package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/itsLeonB/ezutil"
	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/logging"
	"github.com/itsLeonB/together-list/internal/service"
	"github.com/itsLeonB/together-list/internal/util"
	"github.com/rotisserie/eris"
)

type MessageHandler struct {
	listService *service.ListService
}

func NewMessageHandler(listService *service.ListService) *MessageHandler {
	return &MessageHandler{
		listService: listService,
	}
}

func (mh *MessageHandler) HandleSave() bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		message := update.Message
		chatID := message.Chat.ID
		messageID := message.ID

		header, body := util.SplitFirstLine(message.Text)
		if strings.TrimSpace(body) == "" {
			sendMessage(b, ctx, chatID, appconstant.NoURL, messageID)
			return
		}

		logging.Infof("Handling message from: %s. Full text: %s", message.From.Username, message.Text)

		statusChan := make(chan string)

		// start status update responder
		go sendStatusUpdates(b, statusChan, chatID, messageID)

		// process the message
		responses, errs := mh.listService.SaveMessage(ctx, header, body, statusChan)

		close(statusChan)

		// collect messages to respond with
		allMessages := aggregateMessages(responses, errs)

		// send final response
		sendMessage(b, ctx, chatID, strings.Join(allMessages, "\n\n"), messageID)
	}
}

func (mh *MessageHandler) HandleHelp() bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		message := update.Message

		helpText := mh.listService.GetHelpString()
		if helpText == "" {
			helpText = "No help information available."
		}

		sendMessage(b, ctx, message.Chat.ID, helpText, message.ID)
	}
}

func sendStatusUpdates(b *bot.Bot, statusChan <-chan string, chatID int64, messageID int) {
	for msg := range statusChan {
		if msg == "" {
			continue
		}
		sendMessage(b, context.Background(), chatID, msg, messageID)
	}
}

func sendMessage(b *bot.Bot, ctx context.Context, chatID int64, msg string, messageID int) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   msg,
		ReplyParameters: &models.ReplyParameters{
			MessageID: messageID,
		},
	})
	if err != nil {
		logging.Errorf("Error replying: %s", err.Error())
	}
}

func aggregateMessages(responses []string, errs []error) []string {
	var messages []string

	for _, res := range responses {
		if strings.TrimSpace(res) != "" {
			messages = append(messages, res)
		}
	}

	for _, err := range errs {
		switch e := err.(type) {
		case ezutil.AppError:
			logging.Errorf("%s", e.Details)
			messages = append(messages, fmt.Sprintf("%s", e.Details))
		default:
			logging.Errorf("Unexpected error: %s", eris.ToString(err, true))
			messages = append(messages, "Unexpected error occurred")
		}
	}

	return messages
}
