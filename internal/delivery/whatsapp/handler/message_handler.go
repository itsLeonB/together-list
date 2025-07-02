package handler

import (
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type MessageHandler struct {
	logger waLog.Logger
}

func NewMessageHandler(logger waLog.Logger) *MessageHandler {
	return &MessageHandler{logger: logger}
}

func (mh *MessageHandler) HandleMessage() func(event any) {
	return func(event any) {
		switch v := event.(type) {
		case *events.Message:
			if v.Info.IsFromMe {
				return
			}
			mh.logger.Infof("Received a message! %s", v.Message.GetExtendedTextMessage().GetText())
		}
	}
}
