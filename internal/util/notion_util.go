package util

import (
	"github.com/itsLeonB/together-list/internal/appconstant"
	"github.com/jomei/notionapi"
	"github.com/rotisserie/eris"
)

func IsTitlePending(page notionapi.Page) (bool, error) {
	title, ok := page.Properties["title"]
	if !ok {
		return false, eris.New("title does not exist")
	}

	titleProp, ok := title.(*notionapi.TitleProperty)
	if !ok {
		return false, eris.New("title is not title property")
	}

	if len(titleProp.Title) < 1 {
		return false, eris.New("rich text is empty")
	}

	text := titleProp.Title[0].Text
	if text == nil {
		return false, eris.New("text is nil")
	}

	return text.Content == appconstant.PendingTitle, nil
}

func GetExtractedLink(page notionapi.Page) (string, error) {
	extractedLinkField, ok := page.Properties["extractedLink"]
	if !ok {
		return "", eris.New("extractedLink does not exist")
	}

	urlProp, ok := extractedLinkField.(*notionapi.URLProperty)
	if !ok {
		return "", eris.New("extractedLink is not a URL property")
	}

	return urlProp.URL, nil
}
