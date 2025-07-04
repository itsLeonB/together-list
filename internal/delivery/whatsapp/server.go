package whatsapp

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/whatsapp/provider"
	"github.com/itsLeonB/together-list/internal/delivery/worker"
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

	var w *worker.Worker
	if configs.AttachWorker {
		log.Println("starting worker...")
		w = worker.SetupWorker(configs, providers)
		go w.RunAll()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	if w != nil {
		log.Println("stopping worker...")
		w.Stop()
	}

	client.Disconnect()
}
