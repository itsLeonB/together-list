package job

import (
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/together-list/internal/config"
	jobrunners "github.com/itsLeonB/together-list/internal/delivery/job/job_runners"
)

func NewJob(logger ezutil.Logger, cfg config.Config) *ezutil.Job {
	switch cfg.JobName {
	case "Summarize":
		return jobrunners.NewSummarizeJob(logger, cfg)
	default:
		logger.Fatalf("job name: %s does not exist", cfg.JobName)
		return nil
	}
}
