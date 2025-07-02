package util

import (
	"regexp"
	"strings"
)

func ExtractUrls(text string) []string {
	re := regexp.MustCompile(`https?://[^\s]+|www\.[^\s]+|\b[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}\b`)
	candidates := re.FindAllStringIndex(text, -1)

	var results []string
	for _, loc := range candidates {
		start := loc[0]
		end := loc[1]
		match := text[start:end]

		// Skip if part of email address (preceded by '@')
		if start > 0 && text[start-1] == '@' {
			continue
		}

		// Trim common trailing punctuation
		match = strings.TrimRight(match, ".,;:)]")

		results = append(results, match)
	}

	return results
}
