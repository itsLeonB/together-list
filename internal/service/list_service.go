package service

import (
	"context"
	"fmt"
	"strings"
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
	notionRepoRegistry map[string]repository.NotionRepository
	llmService         llm.LLMService
	webScraperService  scrape.WebScraperService
}

func NewListService(
	notionRepos []repository.NotionRepository,
	llmService llm.LLMService,
	webScraperService scrape.WebScraperService,
) *ListService {
	notionRepoRegistry := make(map[string]repository.NotionRepository, len(notionRepos))
	for _, repo := range notionRepos {
		if repo == nil {
			continue
		}
		keyword := repo.GetKeyword()
		if keyword == "" {
			continue
		}
		if _, exists := notionRepoRegistry[keyword]; exists {
			continue // Skip if keyword already exists
		}
		notionRepoRegistry[keyword] = repo
	}

	return &ListService{
		notionRepoRegistry,
		llmService,
		webScraperService,
	}
}

func (ls *ListService) IsKeywordSupported(keyword string) bool {
	_, ok := ls.notionRepoRegistry[keyword]
	return ok
}

func (ls *ListService) GetHelpString() string {
	keywords := make([]string, 0, len(ls.notionRepoRegistry))
	for keyword := range ls.notionRepoRegistry {
		keywords = append(keywords, keyword)
	}

	return fmt.Sprintf(appconstant.HelpText, strings.Join(keywords, ", "))
}

func (ls *ListService) SaveMessage(ctx context.Context, msgType, message string, status chan<- string) ([]string, []error) {
	urls := util.ExtractUrls(message)
	if len(urls) == 0 {
		return nil, []error{ezutil.BadRequestError(appconstant.NoURL)}
	}

	if len(urls) == 1 {
		response, err := ls.saveSingleEntry(ctx, entity.NewDatabasePageEntry{
			URL:             urls[0],
			OriginalMessage: message,
			Type:            msgType,
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
			response, err := ls.saveSingleEntry(ctx, entity.NewDatabasePageEntry{
				URL:             inputUrl,
				OriginalMessage: message,
				Type:            msgType,
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

func (ls *ListService) saveSingleEntry(ctx context.Context, entry entity.NewDatabasePageEntry) (string, error) {
	notionRepo, ok := ls.notionRepoRegistry[entry.Type]
	if !ok {
		return "", eris.Errorf("notion repository not found for type: %s", entry.Type)
	}
	if notionRepo == nil {
		return "", eris.New("notion repository is nil")
	}

	existingPages, err := notionRepo.FindAllByURL(ctx, entry.URL)
	if err != nil {
		return "", err
	}
	if len(existingPages) > 0 {
		pageLinks := make([]string, len(existingPages))
		for i, page := range existingPages {
			pageLinks[i] = page.URL
		}
		return fmt.Sprintf(appconstant.URLAlreadyExists, entry.URL, strings.Join(pageLinks, ", ")), nil
	}

	page, err := notionRepo.AddPage(ctx, entry)
	if err != nil {
		return "", err
	}
	if page == nil {
		return "", eris.New("inserted page is nil")
	}

	return fmt.Sprintf(appconstant.MessageSaved, page.URL), nil
}

func (ls *ListService) SummarizeEntry(ctx context.Context) error {
	for _, notionRepo := range ls.notionRepoRegistry {
		ok, err := ls.trySummarizeEntry(ctx, notionRepo)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
	}

	return nil
}

func (ls *ListService) trySummarizeEntry(ctx context.Context, notionRepo repository.NotionRepository) (bool, error) {
	if notionRepo == nil {
		return false, nil
	}

	page, err := notionRepo.GetSinglePendingPage(ctx)
	if err != nil {
		return false, err
	}
	if page.ID == "" {
		return false, nil
	}

	isPending, err := util.IsTitlePending(page)
	if err != nil {
		return false, err
	}
	if !isPending {
		return false, eris.New("page title is not pending")
	}
	extractedLink, err := util.GetExtractedLink(page)
	if err != nil {
		return false, err
	}
	if extractedLink == "" {
		return false, eris.New("extractedLink is empty")
	}

	html, err := ls.webScraperService.GetHTML(extractedLink)
	if err != nil {
		return false, err
	}

	prompt := fmt.Sprintf(appconstant.PromptSummarizePage, html)

	response, err := ls.llmService.GetResponse(ctx, prompt)
	if err != nil {
		return false, err
	}

	summary, err := util.UnmarshalJSONBlock[dto.PageSummary](response)
	if err != nil {
		return false, err
	}

	summary.PageID = notionapi.PageID(page.ID)

	_, err = notionRepo.UpdatePageSummary(ctx, summary)

	return err == nil, err
}
