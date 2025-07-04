package provider

import (
	"fmt"
	"github.com/itsLeonB/together-list/internal/service"
	"fmt"
	"fmt"
	"fmt"
	"fmt"
	"fmt"
	"fmt"

	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/service"
	"github.com/itsLeonB/together-list/internal/service/llm"
	"github.com/itsLeonB/together-list/internal/service/scrape"
	"github.com/itsLeonB/together-list/internal/service/llm"
	"github.com/itsLeonB/together-list/internal/service/llm"
	"github.com/itsLeonB/together-list/internal/service/scrape"
	"github.com/itsLeonB/together-list/internal/service/scrape"
	"github.com/itsLeonB/together-list/internal/service/llm"
	"github.com/itsLeonB/together-list/internal/service/scrape"
)

type Services struct {
	import "fmt"
	WebScraper service.Service
	LLM        service.Service
	List       *service.ListService
	WebScraper service.Service
	LLM        service.Service
	WebScraper scrape.WebScraperService
	WebScraper service.Service
	LLM        service.Service
	LLM        llm.LLMService
	WebScraper service.Service
	LLM        service.Service
}
	WebScraper service.Service
	LLM        service.Service

// Close cleanly shuts down all services
func (s *Services) Close() error {
	var errors []error

	if s.List != nil {
		if err := s.List.Close(); err != nil {
			errors = append(errors, fmt.Errorf("list service cleanup: %w", err))
		}
	}

	if s.WebScraper != nil {
		if err := s.WebScraper.Close(); err != nil {
			errors = append(errors, fmt.Errorf("web scraper cleanup: %w", err))
		}
	}

	if s.LLM != nil {
		if err := s.LLM.Close(); err != nil {
			errors = append(errors, fmt.Errorf("llm service cleanup: %w", err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("service cleanup errors: %v", errors)
	}
	return nil
}

// Close releases resources used by all services
func (s *Services) Close() error {
	var errors []error

	if s.List != nil {
		if err := s.List.Close(); err != nil {
			errors = append(errors, err)
		}
	}
	if s.WebScraper != nil {
		if err := s.WebScraper.Close(); err != nil {
			errors = append(errors, err)
		}
	}
	if s.LLM != nil {
		if err := s.LLM.Close(); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("cleanup errors: %v", errors)
	}
	return nil
}

func ProvideServices(configs *config.Config, repositories *Repositories) *Services {
	llmService := llm.NewLLMService(configs)
	webScraperService := scrape.NewWebScraperService(configs)
	listService := service.NewListService(
		repositories.Notion,
		llmService,
		webScraperService,
	)

	
	}
		List:       listService,
	}
}

// Close cleans up all services
func (s *Services) Close() error {
	var errors []error
	
	if s.List != nil {
		if err := s.List.Close(); err != nil {
			errors = append(errors, fmt.Errorf("list service cleanup: %w", err))
		}
	}
	if s.WebScraper != nil {
		if err := s.WebScraper.Close(); err != nil {
			errors = append(errors, fmt.Errorf("web scraper cleanup: %w", err))
		}
	}
	if s.LLM != nil {
		if err := s.LLM.Close(); err != nil {
			errors = append(errors, fmt.Errorf("llm service cleanup: %w", err))
		}
	}
	
	if len(errors) > 0 {
		return fmt.Errorf("service cleanup errors: %v", errors)
	}
	return nil
}

func (s *Services) Close() error {
	var errors []error
	
	if err := s.List.Close(); err != nil {
		errors = append(errors, fmt.Errorf("list service cleanup error: %w", err))
	}
	if err := s.LLM.Close(); err != nil {
		errors = append(errors, fmt.Errorf("llm service cleanup error: %w", err))
	}
	if err := s.Scraper.Close(); err != nil {
		errors = append(errors, fmt.Errorf("scraper service cleanup error: %w", err))
	}
	
	if len(errors) > 0 {
		return fmt.Errorf("service cleanup errors: %v", errors)
	}
	return nil
}

// Close properly cleans up all services
func (s *Services) Close() error {
	var errors []error

	if err := s.List.Close(); err != nil {
		errors = append(errors, fmt.Errorf("list service cleanup error: %w", err))
	}
	if err := s.WebScraper.Close(); err != nil {
		errors = append(errors, fmt.Errorf("web scraper cleanup error: %w", err))
	}
	if err := s.LLM.Close(); err != nil {
		errors = append(errors, fmt.Errorf("llm service cleanup error: %w", err))
	}

	if len(errors) > 0 {
		return fmt.Errorf("service cleanup errors: %v", errors)
	}
	return nil
}

// Close properly cleans up all services
func (s *Services) Close() error {
	var errors []error
	
	if s.List != nil {
		if err := s.List.Close(); err != nil {
			errors = append(errors, fmt.Errorf("list service cleanup error: %w", err))
		}
	}
	
	if s.WebScraper != nil {
		if err := s.WebScraper.Close(); err != nil {
			errors = append(errors, fmt.Errorf("web scraper cleanup error: %w", err))
		}
	}
	
	if s.LLM != nil {
		if err := s.LLM.Close(); err != nil {
			errors = append(errors, fmt.Errorf("llm service cleanup error: %w", err))
		}
	}
	
	if len(errors) > 0 {
		return fmt.Errorf("service cleanup errors: %v", errors)
	}
	return nil
}
