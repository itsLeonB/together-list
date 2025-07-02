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

	t.Run("extracts URLs without protocol", func(t *testing.T) {
		input := "Check www.example.com and also https://secure.com"

		result := util.ExtractUrls(input)

		expected := []string{
			"www.example.com",
			"https://secure.com",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("extracts bare domain names like google.com", func(t *testing.T) {
		input := "Some links: google.com, www.test.org, and https://secure.net"

		result := util.ExtractUrls(input)

		expected := []string{
			"google.com",
			"www.test.org",
			"https://secure.net",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("does not extract domain from email", func(t *testing.T) {
		input := "Contact us at support@example.com or visit example.com"
		result := util.ExtractUrls(input)

		expected := []string{"example.com"}
		assert.Equal(t, expected, result)
	})
}
