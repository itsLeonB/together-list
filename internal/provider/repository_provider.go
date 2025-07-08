package provider

import (
	"context"

	"github.com/itsLeonB/ezutil"
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/logging"
	"github.com/itsLeonB/together-list/internal/repository"
)

type Repositories struct {
	Notion []repository.NotionRepository
}

func ProvideRepositories(configs *config.Config) *Repositories {
	postgresRepository := repository.NewPostgresRepository(configs.DatabaseUrl)
	defer func() {
		if err := postgresRepository.Close(); err != nil {
			logging.Errorf("failed to close database connection: %v", err)
		}
	}()

	notionDBs, err := postgresRepository.GetNotionDatabases(context.Background())
	if err != nil {
		logging.Fatalf("Failed to get database ID map: %v", err)
	}

	return &Repositories{
		Notion: ezutil.MapSlice(notionDBs, repository.NewNotionRepository),
	}
}
