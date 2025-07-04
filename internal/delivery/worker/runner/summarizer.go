package runner

import (
	"context"
	"log"

	"github.com/go-co-op/gocron"
	"github.com/itsLeonB/together-list/internal/service"
	"github.com/itsLeonB/together-list/internal/util"
)

type summarizer struct {
	name        string
	scheduler   *gocron.Scheduler
	listService *service.ListService
}

func newSummarizer(
	scheduler *gocron.Scheduler,
	listService *service.ListService,
) Runner {
	return &summarizer{
		"Summarizer",
		scheduler,
		listService,
	}
}

func (s *summarizer) Run() {
	_, err := s.scheduler.Every(1).Day().Do(s.summarize)
	if err != nil {
		log.Fatalf("[%s] error creating job: %s\n", s.name, err.Error())
	} else {
		log.Printf("[%s] job successfully created\n", s.name)
	}
}

func (s *summarizer) summarize() {
	var err error
	latency := util.MeasureLatency(func() {
		err = s.listService.SummarizeEntry(context.Background())
	})
	if err != nil {
		log.Printf("[%s] error summarizing page: %s\n", s.name, err.Error())
	} else {
		log.Printf("[%s] success summarizing page", s.name)
		log.Printf("[%s] latency: %d ms", s.name, latency.Milliseconds())
	}
}
