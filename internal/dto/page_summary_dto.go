package dto

import "github.com/jomei/notionapi"

type PageSummary struct {
	PageID  notionapi.PageID
	Title   string `json:"title"`
	Summary string `json:"summary"`
}
