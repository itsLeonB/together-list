package llm

import (
	"context"

	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/logging"
)

type LLMService interface {
	GetResponse(ctx context.Context, prompt string) (string, error)
}

func NewLLMService(configs *config.Config) LLMService {
	if len(configs.LlmProviders) > 0 {
		return newFallbackLLMService(configs)
	}

	return newSingleLLMService(configs.LlmProvider, configs)
}

func newSingleLLMService(provider string, configs *config.Config) LLMService {
	switch provider {
	case appconstant.GoogleLLM:
		return newGoogleLLMService(configs.GoogleLlmApiKey, configs.GoogleLlmModel)
	case appconstant.OpenRouter:
		return newOpenRouterService(configs.OpenRouterApiKey, configs.OpenRouterModel)
	default:
		logging.Warn("no LLM provider configured")
		return nil
	}
}
