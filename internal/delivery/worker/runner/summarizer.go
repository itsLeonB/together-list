package runner

import (
	"context"
	"log"

	"github.com/go-co-op/gocron"
	"github.com/itsLeonB/together-list/internal/service"
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
	if err := s.listService.SummarizeEntry(context.Background()); err != nil {
		log.Printf("[%s] ERROR summarizing: %s\n", s.name, err.Error())
	} else {
		log.Printf("[%s] SUCCESS summarizing page", s.name)
	}
}
