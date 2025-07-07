package job

import (
	"os"

	"github.com/itsLeonB/together-list/internal/config"
	jobrunner "github.com/itsLeonB/together-list/internal/delivery/job/job_runner"
	"github.com/itsLeonB/together-list/internal/logging"
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
	logging.Infof("[%s] setting up job...", j.jobName)
	if err := j.runner.Setup(configs); err != nil {
		logging.Errorf("[%s] error setting up job: %v", j.jobName, err)
		os.Exit(1)
	}

	logging.Infof("[%s] running job...", j.jobName)

	var jobErr error

	latency := util.MeasureLatency(func() { jobErr = j.runner.Run() })

	if jobErr != nil {
		logging.Errorf("[%s] error running job: %v", j.jobName, jobErr)
		os.Exit(1)
	}

	logging.Infof("[%s] success running job for %d ms", j.jobName, latency.Milliseconds())
}
