package repository

import (
	"context"

	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/dto"
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

func (nr *NotionRepository) AddPage(ctx context.Context, entry entity.DatabasePageEntry) (*notionapi.Page, error) {
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

func (nr *NotionRepository) GetSinglePendingPage(ctx context.Context) (notionapi.Page, error) {
	response, err := nr.client.Database.Query(ctx, notionapi.DatabaseID(nr.databaseID), &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: "title",
			RichText: &notionapi.TextFilterCondition{
				Equals: appconstant.PendingTitle,
			},
		},
		PageSize: 1,
	})
	if err != nil {
		return notionapi.Page{}, eris.Wrap(err, "failed to query pages")
	}

	if len(response.Results) == 0 {
		return notionapi.Page{}, nil
	}

	return response.Results[0], nil
}

func (nr *NotionRepository) UpdatePageSummary(ctx context.Context, summary dto.PageSummary) (notionapi.Page, error) {
	response, err := nr.client.Page.Update(ctx, summary.PageID, &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"title": notionapi.TitleProperty{
				Type: notionapi.PropertyTypeTitle,
				Title: []notionapi.RichText{
					{
						Type: notionapi.ObjectTypeText,
						Text: &notionapi.Text{
							Content: summary.Title,
						},
					},
				},
			},
			"summary": notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{
					{
						Type: notionapi.ObjectTypeText,
						Text: &notionapi.Text{
							Content: summary.Summary,
						},
					},
				},
			},
		},
	})

	if err != nil {
		return notionapi.Page{}, eris.Wrap(err, "error updating page")
	}

	return *response, nil
}
