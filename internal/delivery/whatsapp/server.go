package whatsapp

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/whatsapp/provider"
	internalProvider "github.com/itsLeonB/together-list/internal/provider"
)

func Run(configs *config.Config, providers *internalProvider.Providers) {
	loggers := provider.ProvideLoggers(configs)
	stores := provider.ProvideStores(configs, loggers)
	client := SetupClient(loggers, stores)

	if providers == nil {
		providers = internalProvider.ProvideAll(configs)
	}
	handlers := provider.ProvideHandlers(configs, loggers, client, providers.Services)

	SetupHandlers(client, handlers)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
