package llm

import "github.com/itsLeonB/together-list/internal/service"

import "github.com/itsLeonB/together-list/internal/service"

import (
	"github.com/itsLeonB/together-list/internal/service"
	"github.com/itsLeonB/together-list/internal/service"
	"fmt"
	"context"

	"github.com/itsLeonB/together-list/internal/service"
)

type LLMService interface {
	service.Service
	service.Service
	Service
	Service
	Service
	service.Service
	GetResponse(ctx context.Context, prompt string) (string, error)
}

// llmServiceWrapper wraps any LLM implementation with Close method
type llmServiceWrapper struct {
	LLMService
}

func (w *llmServiceWrapper) Close() error {
	return nil
}

// wrapLLMService wraps an LLM service to implement the Service interface
func wrapLLMService(service LLMService) LLMService {
	return &llmServiceWrapper{LLMService: service}
}
