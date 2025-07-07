package jobrunner

import "github.com/itsLeonB/together-list/internal/config"

type JobRunner interface {
	Setup(configs *config.Config) error
	Run() error
}

func NewJobRunner(name string) JobRunner {
	switch name {
	case "Summarize":
		return NewSummarizeJob()
	default:
		return nil
	}
}
