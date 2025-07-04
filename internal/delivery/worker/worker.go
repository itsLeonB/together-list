package worker

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/worker/runner"
	"github.com/itsLeonB/together-list/internal/provider"
)

type Worker struct {
	scheduler *gocron.Scheduler
	runners   []runner.Runner
}

func SetupWorker(configs *config.Config, providers *provider.Providers) *Worker {
	timezone, err := time.LoadLocation(configs.Timezone)
	if err != nil {
		log.Fatalf("error loading time location: %s", err.Error())
	}

	scheduler := gocron.NewScheduler(timezone)

	if providers == nil {
		providers = provider.ProvideAll(configs)
	}

	runners := runner.ProvideRunners(scheduler, providers.Services)

	return &Worker{
		scheduler,
		runners,
	}
}

func (w *Worker) RunAll() {
	log.Println("running runners...")

	for _, runner := range w.runners {
		runner.Run()
	}

	w.scheduler.StartAsync()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	log.Println("shutting down worker...")

	w.scheduler.Stop()

	log.Println("worker successfully shutdown")
}

func (w *Worker) Stop() {
	if w.scheduler != nil {
		w.scheduler.Stop()
	}
	log.Println("worker stopped")
}
