package util

import (
	"encoding/json"
	"strings"

	"github.com/rotisserie/eris"
)

func SplitFirstLine(inputText string) (string, string) {
	newlineIndex := strings.Index(inputText, "\n")

	if newlineIndex == -1 {
		return inputText, ""
	}

	return inputText[:newlineIndex], inputText[newlineIndex+1:]
}

func UnmarshalJSONBlock[T any](input string) (T, error) {
	var result T

	// Trim surrounding whitespace
	clean := strings.TrimSpace(input)

	// Unwrap triple backticks with or without json marker
	clean = strings.TrimPrefix(clean, "```json")
	clean = strings.TrimPrefix(clean, "```")
	clean = strings.TrimSuffix(clean, "```")
	clean = strings.TrimSpace(clean)

	if clean == "" {
		return result, eris.New("empty JSON after unwrapping")
	}

	// Unmarshal to generic type
	if err := json.Unmarshal([]byte(clean), &result); err != nil {
		return result, eris.Wrapf(err, "error unmarshaling json: %s into %T", clean, result)
	}

	return result, nil
}
