package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/itsLeonB/ezutil"
	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/dto"
	"github.com/itsLeonB/together-list/internal/entity"
	"github.com/itsLeonB/together-list/internal/repository"
	"github.com/itsLeonB/together-list/internal/service/llm"
	"github.com/itsLeonB/together-list/internal/service/scrape"
	"github.com/itsLeonB/together-list/internal/util"
	"github.com/jomei/notionapi"
	"github.com/rotisserie/eris"
)

type ListService struct {
	notionRepository  *repository.NotionRepository
	llmService        llm.LLMService
	webScraperService scrape.WebScraperService
}

func NewListService(
	notionRepository *repository.NotionRepository,
	llmService llm.LLMService,
	webScraperService scrape.WebScraperService,
) *ListService {
	return &ListService{
		notionRepository,
		llmService,
		webScraperService,
	}
}
func (ls *ListService) SaveMessage(ctx context.Context, message string, status chan<- string) ([]string, []error) {
	urls := util.ExtractUrls(message)
	if len(urls) == 0 {
		return nil, []error{ezutil.BadRequestError(appconstant.NoURL)}
	}

	if len(urls) == 1 {
		response, err := ls.saveSingleEntry(ctx, entity.DatabasePageEntry{
			Title:           "pending",
			URL:             urls[0],
			OriginalMessage: message,
		})
		if err != nil {
			return nil, []error{err}
		}
		return []string{response}, nil
	}

	status <- "Detected multiple URLs..."

	var (
		responses []string
		errors    []error
		wg        sync.WaitGroup
		mu        sync.Mutex
	)

	for _, url := range urls {
		wg.Add(1)
		go func(inputUrl string) {
			defer wg.Done()
			response, err := ls.saveSingleEntry(ctx, entity.DatabasePageEntry{
				Title:           appconstant.PendingTitle,
				URL:             inputUrl,
				OriginalMessage: message,
			})
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errors = append(errors, err)
			} else if response != "" {
				responses = append(responses, response)
			}
		}(url)
	}

	wg.Wait()

	return responses, errors
}

func (ls *ListService) saveSingleEntry(ctx context.Context, entry entity.DatabasePageEntry) (string, error) {
	page, err := ls.notionRepository.AddPage(ctx, entry)
	if err != nil {
		return "", err
	}
	if page == nil {
		return "", eris.New("inserted page is nil")
	}

	return fmt.Sprintf(appconstant.MessageSaved, page.URL), nil
}

func (ls *ListService) SummarizeEntry(ctx context.Context) error {
	pages, err := ls.notionRepository.GetSinglePendingPage(ctx)
	if err != nil {
		return err
	}
	if len(pages) < 1 {
		return nil
	}

	page := pages[0]

	isPending, err := util.IsTitlePending(page)
	if err != nil {
		return err
	}
	if !isPending {
		return eris.New("page title is not pending")
	}
	extractedLink, err := util.GetExtractedLink(page)
	if err != nil {
		return err
	}
	if extractedLink == "" {
		return eris.New("extractedLink is empty")
	}

	html, err := ls.webScraperService.GetHTML(extractedLink)
	if err != nil {
		return err
	}

	prompt := fmt.Sprintf(appconstant.PromptSummarizePage, html)

	response, err := ls.llmService.GetResponse(ctx, prompt)
	if err != nil {
		return err
	}

	summary, err := util.UnmarshalJSONBlock[dto.PageSummary](response)
	if err != nil {
		return err
	}

	summary.PageID = notionapi.PageID(page.ID)

	_, err = ls.notionRepository.UpdatePageSummary(ctx, summary)

	return err
}
