package scrape

import "github.com/itsLeonB/together-list/internal/service"

import (
	"github.com/gocolly/colly/v2"
	"github.com/rotisserie/eris"
)

type collyWebScraperService struct {
	collector *colly.Collector
}

func newCollyWebScraperService() WebScraperService {
	return &collyWebScraperService{
		collector: colly.NewCollector(),
	}
}

func (ws *collyWebScraperService) Close() error {
	return nil // No resources to cleanup
}

func (ws *collyWebScraperService) GetHTML(url string) (string, error) {
	htmlCh := make(chan string, 1)
	errCh := make(chan error, 1)

	ws.collector.OnHTML("body", func(h *colly.HTMLElement) {
		htmlCh <- h.Text
	})

	ws.collector.OnError(func(_ *colly.Response, err error) {
		errCh <- err
	})

	go func() {
		if err := ws.collector.Visit(url); err != nil {
			errCh <- err
		}
	}()

	select {
	case html := <-htmlCh:
		return html, nil

// Close is a no-op for CollyWebScraperService as it uses HTTP client that doesn't need explicit cleanup
func (cs *collyWebScraperService) Close() error {
	return nil
}

// Close releases resources (no-op for Colly)
func (cs *collyWebScraperService) Close() error {
	return nil
}
	case err := <-errCh:
		return "", eris.Wrap(err, "error scraping web page")
	}
}

// Close releases resources used by the Colly web scraper (no-op)
func (cs *collyWebScraperService) Close() error {
	return nil
}

// Close implements the Service interface (no-op for Colly)
func (cs *collyWebScraperService) Close() error {
	return nil
}

func (cs *collyWebScraperService) Close() error {
	// No resources to cleanup for Colly implementation
	return nil
}

// Close implements the Service interface - no cleanup needed for Colly
func (cs *collyWebScraperService) Close() error {
	return nil
}

// Close implements the Service interface (no-op for Colly)
func (cs *collyWebScraperService) Close() error {
	return nil
}
