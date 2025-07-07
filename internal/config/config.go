package config

import (
	"os"

	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/logging"
	"github.com/kelseyhightower/envconfig"
)

type ConfigLoader interface {
	Load() *Config
}

type Config struct {
	Env              string
	DatabaseUrl      string
	MessageKeyword   string
	Timezone         string
	JobName          string
	NotionApiKey     string
	NotionDatabaseId string
	LlmProvider      string
	LlmProviders     []string
	GoogleLlmApiKey  string
	GoogleLlmModel   string
	OpenRouterApiKey string
	OpenRouterModel  string
	WebScraper       string
}

func NewConfigLoader() ConfigLoader {
	serviceType := os.Getenv(appconstant.ServiceTypeEnvKey)
	switch serviceType {
	case appconstant.ServiceWhatsapp:
		var config whatsappConfigLoader
		if err := envconfig.Process("", &config); err != nil {
			logging.Errorf("error loading whatsapp config: %v", err)
			os.Exit(1)
		}
		return &config
	case appconstant.ServiceJob:
		var config jobConfig
		if err := envconfig.Process("", &config); err != nil {
			logging.Errorf("error loading job config: %v", err)
			os.Exit(1)
		}
		return &config
	default:
		logging.Errorf("undefined service type: %s", serviceType)
		os.Exit(1)
		return nil
	}
}
