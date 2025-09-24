package main

import (
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/job"
	"github.com/itsLeonB/together-list/internal/logging"
	"github.com/itsLeonB/together-list/internal/provider"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logging.Init()
	configLoader := config.NewConfigLoader()
	configs := configLoader.Load()
	logger := provider.ProvideLogger(configs.JobName, configs.Env)
	j := job.NewJob(logger, *configs)
	j.Run()
}
