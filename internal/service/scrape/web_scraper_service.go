package scrape

import (
	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/logging"
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
		logging.Warn("no web scraper configured")
		return nil
	}
}
