package provider

import (
	"github.com/itsLeonB/together-list/internal/config"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type Loggers struct {
	DB     waLog.Logger
	Client waLog.Logger
}

func ProvideLoggers(configs *config.Config) *Loggers {
	logLevel := "INFO"
	if configs.Env == "debug" {
		logLevel = "DEBUG"
	}

	return &Loggers{
		DB:     waLog.Stdout("Database", logLevel, true),
		Client: waLog.Stdout("Client", logLevel, true),
	}
}
