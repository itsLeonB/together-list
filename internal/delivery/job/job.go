package job

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/itsLeonB/together-list/internal/config"
	jobrunner "github.com/itsLeonB/together-list/internal/delivery/job/job_runner"
	"github.com/itsLeonB/together-list/internal/util"
)

type Job struct {
	jobName string
	runner  jobrunner.JobRunner
}

func NewJob(jobName string) *Job {
	return &Job{
		jobName,
		jobrunner.NewJobRunner(jobName),
	}
}

func (j *Job) Run(configs *config.Config) {
	slog.Info(fmt.Sprintf("[%s] setting up job...", j.jobName))
	if err := j.runner.Setup(configs); err != nil {
		slog.Error(fmt.Sprintf("[%s] error setting up job: %v", j.jobName, err))
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("[%s] running job...", j.jobName))

	var jobErr error

	latency := util.MeasureLatency(func() { jobErr = j.runner.Run() })

	if jobErr != nil {
		slog.Error(fmt.Sprintf("[%s] error running job: %v", j.jobName, jobErr))
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("[%s] success running job for %d ms", j.jobName, latency.Milliseconds()))
}
