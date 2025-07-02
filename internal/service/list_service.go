package service

import (
	"fmt"
	"sync"

	"github.com/itsLeonB/ezutil"
	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/entity"
	"github.com/itsLeonB/together-list/internal/repository"
	"github.com/itsLeonB/together-list/internal/util"
	"github.com/rotisserie/eris"
)

type ListService struct {
	notionRepository *repository.NotionRepository
}

func NewListService(
	notionRepository *repository.NotionRepository,
) *ListService {
	return &ListService{
		notionRepository,
	}
}
func (ls *ListService) SaveMessage(message string, status chan<- string) ([]string, []error) {
	urls := util.ExtractUrls(message)
	if len(urls) == 0 {
		return nil, []error{ezutil.BadRequestError(appconstant.NoURL)}
	}

	if len(urls) == 1 {
		response, err := ls.saveSingleEntry(entity.DatabasePageEntry{
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
			response, err := ls.saveSingleEntry(entity.DatabasePageEntry{
				Title:           "pending",
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

func (ls *ListService) saveSingleEntry(entry entity.DatabasePageEntry) (string, error) {
	page, err := ls.notionRepository.AddPageToDatabase(entry)
	if err != nil {
		return "", err
	}
	if page == nil {
		return "", eris.New("inserted page is nil")
	}

	return fmt.Sprintf(appconstant.MessageSaved, page.URL), nil
}
