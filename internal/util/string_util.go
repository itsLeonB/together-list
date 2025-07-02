package util

import "strings"

func SplitFirstLine(inputText string) (string, string) {
	newlineIndex := strings.Index(inputText, "\n")

	if newlineIndex == -1 {
		return inputText, ""
	}

	return inputText[:newlineIndex], inputText[newlineIndex+1:]
}
