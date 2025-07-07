package main

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/job"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	configLoader := config.NewConfigLoader()
	configs := configLoader.Load()
	j := job.NewJob(configs.JobName)
	j.Run(configs)
}
