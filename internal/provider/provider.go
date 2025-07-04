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

// Close cleans up all provider resources
func (p *Providers) Close() error {
	if p.Services != nil {
		return p.Services.Close()
	}
	return nil
}

// Close properly cleans up all provider resources
func (p *Providers) Close() error {
	if p.Services != nil {
		return p.Services.Close()
	}
	return nil
}
