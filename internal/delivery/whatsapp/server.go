package whatsapp

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/whatsapp/provider"
)

func Run(configs *config.Config) {
	loggers := provider.ProvideLoggers(configs)
	handlers := provider.ProvideHandlers(loggers)
	stores := provider.ProvideStores(configs, loggers)

	client := SetupClient(handlers, loggers, stores)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
