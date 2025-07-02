package main

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/whatsapp"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := config.LoadConfig()
	whatsapp.Run(config)
}
