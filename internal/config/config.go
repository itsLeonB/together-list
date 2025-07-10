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
	JobName          string
	LlmProvider      string
	LlmProviders     []string
	GoogleLlmApiKey  string
	GoogleLlmModel   string
	OpenRouterApiKey string
	OpenRouterModel  string
	WebScraper       string
	TelegramBotToken string
}

func NewConfigLoader() ConfigLoader {
	serviceType := os.Getenv(appconstant.ServiceTypeEnvKey)
	configLoader := getConfigLoader(serviceType)
	if configLoader == nil {
		logging.Fatalf("undefined service type: %s", serviceType)
		return nil
	}
	if err := envconfig.Process("", configLoader); err != nil {
		logging.Fatalf("error loading %s config: %v", serviceType, err)
	}
	return configLoader
}

func getConfigLoader(serviceType string) ConfigLoader {
	switch serviceType {
	case appconstant.ServiceWhatsapp:
		return &whatsappConfigLoader{}
	case appconstant.ServiceJob:
		return &jobConfigLoader{}
	case appconstant.ServiceTelegram:
		return &telegramConfigLoader{}
	default:
		return nil
	}
}
