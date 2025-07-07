package jobrunner

import (
	"os"

	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/logging"
)

type JobRunner interface {
	Setup(configs *config.Config) error
	Run() error
}

func NewJobRunner(name string) JobRunner {
	switch name {
	case "Summarize":
		return NewSummarizeJob()
	default:
		logging.Errorf("job name: %s does not exist", name)
		os.Exit(1)
		return nil
	}
}
