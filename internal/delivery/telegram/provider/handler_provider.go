package provider

import (
	"github.com/itsLeonB/together-list/internal/delivery/telegram/handler"
	"github.com/itsLeonB/together-list/internal/provider"
)

type Handlers struct {
	Message *handler.MessageHandler
}

func ProvideHandlers(
	services *provider.Services,
) *Handlers {
	return &Handlers{
		Message: handler.NewMessageHandler(services.List),
	}
}
