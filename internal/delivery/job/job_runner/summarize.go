package jobrunner

import (
	"context"

	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/provider"
	"github.com/itsLeonB/together-list/internal/service"
)

type summarize struct {
	listService *service.ListService
}

func NewSummarizeJob() JobRunner {
	return &summarize{}
}

func (s *summarize) Setup(configs *config.Config) error {
	providers := provider.ProvideAll(configs)
	s.listService = providers.Services.List
	return nil
}

func (s *summarize) Run() error {
	return s.listService.SummarizeEntry(context.Background())
}
