package provider

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/service"
	"github.com/itsLeonB/together-list/internal/service/llm"
	"github.com/itsLeonB/together-list/internal/service/scrape"
)

type Services struct {
	List *service.ListService
}

func ProvideServices(configs *config.Config, repositories *Repositories) *Services {
	llmService := llm.NewLLMService(configs)
	webScraperService := scrape.NewWebScraperService(configs)

	return &Services{
		List: service.NewListService(
			repositories.Notion,
			llmService,
			webScraperService,
		),
	}
}
