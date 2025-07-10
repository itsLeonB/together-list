package telegram

import (
	"github.com/go-telegram/bot"
	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/delivery/telegram/provider"
)

func setupHandlers(b *bot.Bot, handlers *provider.Handlers) {
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		appconstant.HelpKeyword,
		bot.MatchTypePrefix,
		handlers.Message.HandleHelp(),
	)

	for _, keyword := range handlers.Keywords {
		b.RegisterHandler(
			bot.HandlerTypeMessageText,
			keyword,
			bot.MatchTypePrefix,
			handlers.Message.HandleSave(),
		)
	}
}
