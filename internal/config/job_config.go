package config

import (
	"os"

	"github.com/itsLeonB/together-list/internal/logging"
)

type jobConfig struct {
	Env              string   `required:"true"`
	JobName          string   `split_words:"true" required:"true"`
	NotionApiKey     string   `split_words:"true" required:"true"`
	NotionDatabaseId string   `split_words:"true" required:"true"`
	LlmProvider      string   `split_words:"true"`
	LlmProviders     []string `split_words:"true"`
	GoogleLlmApiKey  string   `split_words:"true"`
	GoogleLlmModel   string   `split_words:"true"`
	OpenRouterApiKey string   `split_words:"true"`
	OpenRouterModel  string   `split_words:"true"`
	WebScraper       string   `split_words:"true" required:"true"`
}

func (jc *jobConfig) Load() *Config {
	if jc.LlmProvider == "" && len(jc.LlmProviders) == 0 {
		logging.Error("LLM_PROVIDER or LLM_PROVIDERS must be set")
		os.Exit(1)
		return nil
	}

	return &Config{
		Env:              jc.Env,
		JobName:          jc.JobName,
		NotionApiKey:     jc.NotionApiKey,
		NotionDatabaseId: jc.NotionDatabaseId,
		LlmProvider:      jc.LlmProvider,
		LlmProviders:     jc.LlmProviders,
		GoogleLlmApiKey:  jc.GoogleLlmApiKey,
		GoogleLlmModel:   jc.GoogleLlmModel,
		OpenRouterApiKey: jc.OpenRouterApiKey,
		OpenRouterModel:  jc.OpenRouterModel,
		WebScraper:       jc.WebScraper,
	}
}
