package provider

import "github.com/itsLeonB/together-list/internal/delivery/whatsapp/handler"

type Handlers struct {
	Message *handler.MessageHandler
}

func ProvideHandlers(loggers *Loggers) *Handlers {
	return &Handlers{
		Message: handler.NewMessageHandler(loggers.Client),
	}
}
