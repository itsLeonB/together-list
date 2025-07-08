package repository

import (
	"context"
	"time"

	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/itsLeonB/together-list/internal/dto"
	"github.com/itsLeonB/together-list/internal/entity"
	"github.com/jomei/notionapi"
	"github.com/rotisserie/eris"
)

type notionRepositoryImpl struct {
	keyword    string
	client     *notionapi.Client
	databaseID notionapi.DatabaseID
}

func NewNotionRepository(db entity.NotionDatabase) NotionRepository {
	return &notionRepositoryImpl{
		keyword:    db.Keyword,
		client:     notionapi.NewClient(notionapi.Token(db.APIKey)),
		databaseID: notionapi.DatabaseID(db.DatabaseID),
	}
}

func (nr *notionRepositoryImpl) GetKeyword() string {
	return nr.keyword
}

func (nr *notionRepositoryImpl) AddPage(ctx context.Context, entry entity.NewDatabasePageEntry) (*notionapi.Page, error) {
	newPage, err := nr.client.Page.Create(ctx, &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: nr.databaseID,
		},
		Properties: notionapi.Properties{
			"title": stringToTitleProperty(appconstant.PendingTitle),
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
			"createdAt": timeToDateProperty(time.Now()),
			"updatedAt": timeToDateProperty(time.Now()),
		},
	})

	if err != nil {
		return nil, eris.Wrap(err, "failed to create page")
	}

	return newPage, nil
}

func (nr *notionRepositoryImpl) GetSinglePendingPage(ctx context.Context) (notionapi.Page, error) {
	response, err := nr.client.Database.Query(ctx, nr.databaseID, &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: "title",
			RichText: &notionapi.TextFilterCondition{
				Equals: appconstant.PendingTitle,
			},
		},
		Sorts: []notionapi.SortObject{
			{
				Property:  "createdAt",
				Direction: notionapi.SortOrderDESC,
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

func (nr *notionRepositoryImpl) UpdatePageSummary(ctx context.Context, summary dto.PageSummary) (notionapi.Page, error) {
	response, err := nr.client.Page.Update(ctx, summary.PageID, &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"title": stringToTitleProperty(summary.Title),
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
			"updatedAt": timeToDateProperty(time.Now()),
		},
	})

	if err != nil {
		return notionapi.Page{}, eris.Wrap(err, "error updating page")
	}

	return *response, nil
}

func timeToDateProperty(t time.Time) notionapi.DateProperty {
	startDate := notionapi.Date(t)
	return notionapi.DateProperty{
		Type: notionapi.PropertyTypeDate,
		Date: &notionapi.DateObject{
			Start: &startDate,
		},
	}
}

func stringToTitleProperty(title string) notionapi.TitleProperty {
	return notionapi.TitleProperty{
		Type: notionapi.PropertyTypeTitle,
		Title: []notionapi.RichText{
			{
				Type: notionapi.ObjectTypeText,
				Text: &notionapi.Text{
					Content: title,
				},
			},
		},
	}
}
