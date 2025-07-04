package scrape

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
	case err := <-errCh:
		return "", eris.Wrap(err, "error scraping web page")
	}
}
