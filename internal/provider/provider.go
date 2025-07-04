package provider

import "github.com/itsLeonB/together-list/internal/config"

type Providers struct {
	*Repositories
	*Services
}

func ProvideAll(configs *config.Config) *Providers {
	repositories := ProvideRepositories(configs)
	services := ProvideServices(configs, repositories)

	return &Providers{
		Repositories: repositories,
		Services:     services,
	}
}
