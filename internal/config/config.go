package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env            string `split_words:"true" required:"true"`
	DatabaseUrl    string `split_words:"true" required:"true"`
	MessageKeyword string `split_words:"true" required:"true"`
	Timezone       string `required:"true"`
	AttachWorker   bool   `split_words:"true"`

	NotionApiKey     string `split_words:"true" required:"true"`
	NotionDatabaseId string `split_words:"true" required:"true"`

	LlmProvider      string   `split_words:"true" required:"true"`
	LlmProviders     []string `split_words:"true"`
	GoogleLlmApiKey  string   `split_words:"true" required:"true"`
	GoogleLlmModel   string   `split_words:"true" required:"true"`
	OpenRouterApiKey string   `split_words:"true" required:"true"`
	OpenRouterModel  string   `split_words:"true" required:"true"`

	WebScraper string `split_words:"true" required:"true"`
}

func LoadConfig() *Config {
	var newConfig Config
	if err := envconfig.Process("", &newConfig); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	return &newConfig
}
