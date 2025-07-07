package scrape

import (
	"log/slog"

	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/config"
)

type WebScraperService interface {
	GetHTML(url string) (string, error)
}

func NewWebScraperService(configs *config.Config) WebScraperService {
	switch configs.WebScraper {
	case appconstant.WebScraperColly:
		return newCollyWebScraperService()
	case appconstant.WebScraperChromeDP:
		return newChromeDPWebScraperService()
	default:
		slog.Warn("no web scraper configured")
		return nil
	}
}
