package runner

import (
	"github.com/go-co-op/gocron"
	"github.com/itsLeonB/together-list/internal/provider"
)

type Runner interface {
	Run()
}

func ProvideRunners(
	scheduler *gocron.Scheduler,
	services *provider.Services,
) []Runner {
	return []Runner{
		newSummarizer(scheduler, services.List),
	}
}
