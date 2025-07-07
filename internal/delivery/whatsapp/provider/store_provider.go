package provider

import (
	"context"

	"github.com/itsLeonB/together-list/internal/config"
	"github.com/itsLeonB/together-list/internal/logging"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

type Stores struct {
	Device *store.Device
}

func ProvideStores(configs *config.Config, loggers *Loggers) *Stores {
	ctx := context.Background()
	container, err := sqlstore.New(ctx, "pgx", configs.DatabaseUrl, loggers.DB)
	if err != nil {
		logging.Fatalf("Unable to create database container: %v", err)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		logging.Fatalf("Unable to get device: %v", err)
	}

	return &Stores{
		Device: deviceStore,
	}
}
