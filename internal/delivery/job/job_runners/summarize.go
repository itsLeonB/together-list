package jobrunners

import (
	"context"

	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/provider"
	"github.com/itsLeonB/together-list/internal/service"
)

type summarize struct {
	listService *service.ListService
}

func NewSummarizeJob(logger ezutil.Logger, cfg config.Config) *ezutil.Job {
	providers := provider.ProvideAll(&cfg)
	summarize := summarize{}

	job := ezutil.NewJob(logger, summarize.Run)
	job = job.WithSetupFunc(func() error {
		summarize.listService = providers.List
		return nil
	})

	return job
}

func (s *summarize) Run() error {
	return s.listService.SummarizeEntry(context.Background())
}
