package provider

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/whatsapp/handler"
	"github.com/itsLeonB/together-list/internal/provider"
	"go.mau.fi/whatsmeow"
)

type Handlers struct {
	Message *handler.MessageHandler
}

func ProvideHandlers(
	configs *config.Config,
	loggers *Loggers,
	client *whatsmeow.Client,
	services *provider.Services,
) *Handlers {
	messageHandler := handler.NewMessageHandler(
		configs.MessageKeyword,
		loggers.Client,
		client,
		services.List,
	)

	return &Handlers{
		Message: messageHandler,
	}
}
