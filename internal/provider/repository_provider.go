package provider

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/repository"
)

type Repositories struct {
	Notion *repository.NotionRepository
}

func ProvideRepositories(configs *config.Config) *Repositories {
	return &Repositories{
		Notion: repository.NewNotionRepository(configs.NotionDatabaseId, configs.NotionApiKey),
	}
}
