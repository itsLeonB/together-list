package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/delivery/telegram/provider"
	"github.com/itsLeonB/together-list/internal/logging"
	internalProvider "github.com/itsLeonB/together-list/internal/provider"
)

func Run(configs *config.Config) {
	opts := []bot.Option{}

	b, err := bot.New(configs.TelegramBotToken, opts...)
	if err != nil {
		logging.Fatalf("failed to create telegram bot: %v", err)
	}

	internalProviders := internalProvider.ProvideAll(configs)
	handlers := provider.ProvideHandlers(internalProviders.Services)
	setupHandlers(b, handlers)

	logging.Infof("starting telegram bot...")
	b.Start(context.Background())
}
