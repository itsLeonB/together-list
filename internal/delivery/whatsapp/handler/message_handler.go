package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/itsLeonB/ezutil"
	"github.com/itsLeonB/together-list/internal/service"
	"github.com/itsLeonB/together-list/internal/util"
	"github.com/rotisserie/eris"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type MessageHandler struct {
	messageKeyword string
	logger         waLog.Logger
	client         *whatsmeow.Client
	listService    *service.ListService
}

func NewMessageHandler(
	messageKeyword string,
	logger waLog.Logger,
	client *whatsmeow.Client,
	listService *service.ListService,
) *MessageHandler {
	return &MessageHandler{
		messageKeyword,
		logger,
		client,
		listService,
	}
}

func (mh *MessageHandler) HandleMessage() func(event any) {
	return func(event any) {
		msgEvent, ok := event.(*events.Message)
		if !ok || msgEvent.Info.IsFromMe || msgEvent.Message == nil {
			return
		}

		extMsg := msgEvent.Message.GetExtendedTextMessage()
		if extMsg == nil {
			return
		}

		fullText := extMsg.GetText()
		header, body := util.SplitFirstLine(fullText)

		if header != mh.messageKeyword || strings.TrimSpace(body) == "" {
			return
		}

		ctx := context.Background()
		chatID := msgEvent.Info.Chat
		senderID := msgEvent.Info.Sender
		stanzaID := msgEvent.Info.ID

		ctxInfo := &waE2E.ContextInfo{
			StanzaID:      proto.String(stanzaID),
			Participant:   proto.String(senderID.String()),
			QuotedMessage: msgEvent.Message,
		}

		mh.logger.Infof("Handling message from: %s. Full text: %s", senderID, fullText)

		statusChan := make(chan string)

		// start status update responder
		go mh.sendStatusUpdates(statusChan, chatID, ctxInfo)

		// process the message
		responses, errs := mh.listService.SaveMessage(ctx, body, statusChan)

		close(statusChan)

		// collect messages to respond with
		allMessages := mh.aggregateMessages(responses, errs)

		// send final response
		if err := mh.sendMessage(ctx, chatID, strings.Join(allMessages, "\n\n"), ctxInfo); err != nil {
			mh.logger.Errorf("Error replying: %s", err.Error())
		}
	}
}

func (mh *MessageHandler) sendStatusUpdates(statusChan <-chan string, chatID types.JID, ctxInfo *waE2E.ContextInfo) {
	for msg := range statusChan {
		if msg == "" {
			continue
		}
		if err := mh.sendMessage(context.Background(), chatID, msg, ctxInfo); err != nil {
			mh.logger.Errorf("Error sending status update: %s", err.Error())
		}
	}
}

func (mh *MessageHandler) sendMessage(ctx context.Context, chatID types.JID, msg string, ctxInfo *waE2E.ContextInfo) error {
	_, err := mh.client.SendMessage(ctx, chatID, &waE2E.Message{
		ExtendedTextMessage: &waE2E.ExtendedTextMessage{
			Text:        proto.String(msg),
			ContextInfo: ctxInfo,
		},
	})
	return err
}

func (mh *MessageHandler) aggregateMessages(responses []string, errs []error) []string {
	var messages []string

	for _, res := range responses {
		if strings.TrimSpace(res) != "" {
			messages = append(messages, res)
		}
	}

	for _, err := range errs {
		switch e := err.(type) {
		case ezutil.AppError:
			mh.logger.Errorf("%s", e.Details)
			messages = append(messages, fmt.Sprintf("%s", e.Details))
		default:
			mh.logger.Errorf("Unexpected error: %s", eris.ToString(err, true))
			messages = append(messages, "Unexpected error occurred")
		}
	}

	return messages
}
