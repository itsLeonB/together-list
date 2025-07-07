package main

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/worker"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	configLoader := config.NewConfigLoader()
	configs := configLoader.Load()
	w := worker.SetupWorker(configs, nil)
	w.RunAll()
}
