package llm

import (
	"github.com/itsLeonB/together-list/internal/service"
)

import (
	"github.com/itsLeonB/together-list/internal/service"
	"fmt"
	"fmt"
	"fmt"
	"fmt"
	"context"
	"fmt"

	"github.com/itsLeonB/together-list/internal/config"
	"github.com/rotisserie/eris"
)

type fallbackLLMService struct {
	services []LLMService
}

func newFallbackLLMService(configs *config.Config) LLMService {
	services := make([]LLMService, len(configs.LlmProviders))
	for i, llmProvider := range configs.LlmProviders {
		services[i] = newSingleLLMService(llmProvider, configs)
	}
	return &fallbackLLMService{services: services}
}

func (f *fallbackLLMService) Close() error {
	var errors []error
	for _, service := range f.services {
		if err := service.Close(); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("fallback service cleanup errors: %v", errors)
	}
	return nil
}

func (f *fallbackLLMService) GetResponse(ctx context.Context, prompt string) (string, error) {
	var lastErr error
	for _, service := range f.services {
		response, err := service.GetResponse(ctx, prompt)
		if err == nil {
			return response, nil
		}
		lastErr = err
	}
	return "", eris.Wrap(lastErr, "all LLM services failed")
}

// Close releases resources used by all fallback LLM services
func (f *fallbackLLMService) Close() error {
	var errors []error
	for _, service := range f.services {
		if err := service.Close(); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("fallback service cleanup errors: %v", errors)
	}
	return nil
}

func (f *fallbackLLMService) Close() error {
	var errors []error
	
	for _, service := range f.services {
		if err := service.Close(); err != nil {
			errors = append(errors, err)
		}
	}
	
	if len(errors) > 0 {
		return fmt.Errorf("fallback service cleanup errors: %v", errors)
	}
	return nil
}

// Close properly cleans up all underlying LLM services
func (f *fallbackLLMService) Close() error {
	var errors []error
	for _, service := range f.services {
		if closer, ok := service.(service.Service); ok {
			if err := closer.Close(); err != nil {
				errors = append(errors, err)
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("fallback LLM service cleanup errors: %v", errors)
	}
	return nil
}
