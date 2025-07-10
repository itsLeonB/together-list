package main

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/telegram"
	"github.com/itsLeonB/together-list/internal/logging"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logging.Init()
	configLoader := config.NewConfigLoader()
	configs := configLoader.Load()
	telegram.Run(configs)
}
