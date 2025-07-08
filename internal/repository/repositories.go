package repository

import (
	"context"

	"github.com/itsLeonB/together-list/internal/dto"
	"github.com/itsLeonB/together-list/internal/entity"
	"github.com/jomei/notionapi"
)

type NotionRepository interface {
	GetKeyword() string
	AddPage(ctx context.Context, entry entity.NewDatabasePageEntry) (*notionapi.Page, error)
	GetSinglePendingPage(ctx context.Context) (notionapi.Page, error)
	UpdatePageSummary(ctx context.Context, summary dto.PageSummary) (notionapi.Page, error)
}
