package main

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/worker"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := config.LoadConfig()
	w := worker.SetupWorker(config, nil)
	w.RunAll()
}
