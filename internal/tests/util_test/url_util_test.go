package util_test

import (
	"testing"

	"github.com/itsLeonB/together-list/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestExtractUrls(t *testing.T) {
	t.Run("extracts multiple URLs from multiline text", func(t *testing.T) {
		input := `
http://www.google.com
http://www.notion.com`

		result := util.ExtractUrls(input)

		expected := []string{
			"http://www.google.com",
			"http://www.notion.com",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("returns empty array if no URL is found", func(t *testing.T) {
		input := "no links here"

		result := util.ExtractUrls(input)

		assert.Empty(t, result)
	})

	t.Run("extracts single URL", func(t *testing.T) {
		input := "Visit https://openai.com for info"

		result := util.ExtractUrls(input)

		expected := []string{"https://openai.com"}
		assert.Equal(t, expected, result)
	})
}
