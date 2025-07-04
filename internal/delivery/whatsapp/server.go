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

	// Setup cleanup
	defer func() {
		log.Println("cleaning up services...")
		if err := providers.Services.Close(); err != nil {
			log.Printf("error during service cleanup: %v", err)
		}
	}()

	// Setup cleanup handler
	defer func() {
		if err := providers.Services.Close(); err != nil {
			log.Printf("error during service cleanup: %v", err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Setup cleanup on shutdown
	defer func() {
		log.Println("cleaning up services...")
		if err := providers.Close(); err != nil {
			log.Printf("error during service cleanup: %v", err)
		}
	}()
	<-c

	if w != nil {
		log.Println("stopping worker...")
		// Close worker with cleanup
		if err := w.Close(providers); err != nil {
			log.Printf("error closing worker: %v", err)
		}
	} else {
		// Cleanup services even if no worker
		if err := providers.Services.Close(); err != nil {
			log.Printf("error during service cleanup: %v", err)
		}
	}

	// Cleanup services
	if err := providers.Services.Close(); err != nil {
		log.Printf("error during service cleanup: %v", err)
	}
	// Cleanup services
	log.Println("cleaning up services...")
	if err := providers.Services.Close(); err != nil {
		log.Printf("error during service cleanup: %v", err)
	}
	// Cleanup all services
	if err := providers.Close(); err != nil {
		log.Printf("error during service cleanup: %v", err)
	}
	// Cleanup services
	log.Println("cleaning up services...")
	if err := providers.Services.Close(); err != nil {
		log.Printf("error during service cleanup: %v", err)
	}
	// Cleanup services
	if err := providers.Services.Close(); err != nil {
		log.Printf("error during service cleanup: %v", err)
	}

	// Cleanup services
	log.Println("cleaning up services...")
	if err := providers.Services.Close(); err != nil {
		log.Printf("error during service cleanup: %v", err)
	}
	client.Disconnect()
}
