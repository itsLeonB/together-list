package jobrunner

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/itsLeonB/together-list/internal/config"
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
		slog.Error(fmt.Sprintf("job name: %s does not exist", name))
		os.Exit(1)
		return nil
	}
}
