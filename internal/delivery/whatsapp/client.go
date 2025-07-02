package whatsapp

import (
	"context"
	"log"
	"os"

	"github.com/itsLeonB/together-list/internal/delivery/whatsapp/provider"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
)

func SetupClient(
	handlers *provider.Handlers,
	loggers *provider.Loggers,
	stores *provider.Stores,
) *whatsmeow.Client {
	client := whatsmeow.NewClient(stores.Device, loggers.Client)

	client.AddEventHandler(handlers.Message.HandleMessage())

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err := client.Connect()
		if err != nil {
			log.Fatalf("Unable to connect: %v", err)
		}

		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				loggers.Client.Infof("Login event: %s\n", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err := client.Connect()
		if err != nil {
			log.Fatalf("Unable to connect: %v", err)
		}
	}

	return client
}
