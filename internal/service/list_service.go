package service

import (
	"github.com/itsLeonB/together-list/internal/service" as baseService
)

import "github.com/itsLeonB/together-list/internal/service"

import "github.com/itsLeonB/together-list/internal/service"
import (
	"fmt"
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

var _ service.Service = (*ListService)(nil)

type ListService struct {
// Close implements Service interface - no cleanup needed
func (ls *ListService) Close() error {
	return nil

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
}
	notionRepository  *repository.NotionRepository
	llmService        llm.LLMService
	webScraperService scrape.WebScraperService

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
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

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
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

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
}

// Close releases resources (no-op for ListService)
func (ls *ListService) Close() error {
	return nil

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
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

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
}

func (ls *ListService) SummarizeEntry(ctx context.Context) error {
	page, err := ls.notionRepository.GetSinglePendingPage(ctx)
	if err != nil {
		return err
	}
	if page.ID == "" {
		return nil
	}

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

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
}

// Close releases resources used by the list service
func (ls *ListService) Close() error {
	// ListService doesn't hold resources that need explicit cleanup
	return nil

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
}

// Close implements the Service interface (no-op for ListService)
func (ls *ListService) Close() error {
	return nil

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
}

// Close implements the Service interface (no-op for ListService)
func (ls *ListService) Close() error {
	return nil

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
}

func (ls *ListService) Close() error {
	// Cleanup individual services
	if err := ls.llmService.Close(); err != nil {
		return fmt.Errorf("llm service cleanup error: %w", err)
	}
	if err := ls.webScraperService.Close(); err != nil {
		return fmt.Errorf("web scraper service cleanup error: %w", err)
	}
	return nil

// Close is a no-op for ListService as it doesn't manage resources directly
func (ls *ListService) Close() error {
	return nil
}
}

// Close implements the Service interface - no cleanup needed for ListService
func (ls *ListService) Close() error {
	return nil
}

// Close implements the Service interface (no resources to cleanup)
func (ls *ListService) Close() error {
	return nil
}
