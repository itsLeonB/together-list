package repository

import (
	"context"

	"github.com/itsLeonB/together-list/internal/entity"
	"github.com/jomei/notionapi"
	"github.com/rotisserie/eris"
)

type NotionRepository struct {
	databaseID string
	client     *notionapi.Client
}

func NewNotionRepository(databaseID, token string) *NotionRepository {
	return &NotionRepository{
		databaseID: databaseID,
		client:     notionapi.NewClient(notionapi.Token(token)),
	}
}
func (nr *NotionRepository) AddPageToDatabase(ctx context.Context, entry entity.DatabasePageEntry) (*notionapi.Page, error) {
	newPage, err := nr.client.Page.Create(ctx, &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: notionapi.DatabaseID(nr.databaseID),
		},
		Properties: notionapi.Properties{
			"title": notionapi.TitleProperty{
				Type: notionapi.PropertyTypeTitle,
				Title: []notionapi.RichText{
					{
						Type: notionapi.ObjectTypeText,
						Text: &notionapi.Text{
							Content: entry.Title,
						},
					},
				},
			},
			"extractedLink": notionapi.URLProperty{
				Type: notionapi.PropertyTypeURL,
				URL:  entry.URL,
			},
			"originalMessage": notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{
					{
						Type: notionapi.ObjectTypeText,
						Text: &notionapi.Text{
							Content: entry.OriginalMessage,
						},
					},
				},
			},
		},
	})

	if err != nil {
		return nil, eris.Wrap(err, "failed to create page")
	}

	return newPage, nil
}
