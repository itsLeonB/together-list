package telegram

import (
	"github.com/go-telegram/bot"
	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/delivery/telegram/provider"
)

func setupHandlers(b *bot.Bot, handlers *provider.Handlers) {
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		appconstant.TelegramSaveCommand,
		bot.MatchTypePrefix,
		handlers.Message.HandleSave(),
	)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		appconstant.TelegramHelpCommand,
		bot.MatchTypePrefix,
		handlers.Message.HandleHelp(),
	)
}
