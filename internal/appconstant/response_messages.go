package appconstant

import "fmt"

const (
	NoURL        = "No URL found in the message"
	MultipleURLs = "Multiple URLs found, saving to multiple entries..."
	Error        = "There was an unexpected error. Please contact developer."
	MessageSaved = "Message saved to: %s"
)

func UnsupportedKeyword(keyword string) string {
	return fmt.Sprintf("Unsupported keyword: %s", keyword)
}
