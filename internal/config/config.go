package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env              string `split_words:"true"`
	NotionApiKey     string `split_words:"true"`
	NotionDatabaseId string `split_words:"true"`
	DatabaseUrl      string `split_words:"true"`
	MessageKeyword   string `split_words:"true"`
}

func LoadConfig() *Config {
	var newConfig Config
	if err := envconfig.Process("", &newConfig); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &newConfig
}
