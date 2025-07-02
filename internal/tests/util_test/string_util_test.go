package util_test

import (
	"testing"

	"github.com/itsLeonB/together-list/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestSplitFirstLine(t *testing.T) {
	t.Run("splits first line from the rest", func(t *testing.T) {
		input := "beasiswa\nhttp://www.google.com\nhttp://www.notion.com"
		first, rest := util.SplitFirstLine(input)

		assert.Equal(t, "beasiswa", first)
		assert.Equal(t, "http://www.google.com\nhttp://www.notion.com", rest)
	})

	t.Run("returns entire string as first item if no newline", func(t *testing.T) {
		input := "single line input"
		first, rest := util.SplitFirstLine(input)

		assert.Equal(t, "single line input", first)
		assert.Equal(t, "", rest)
	})

	t.Run("handles empty string", func(t *testing.T) {
		input := ""
		first, rest := util.SplitFirstLine(input)

		assert.Equal(t, "", first)
		assert.Equal(t, "", rest)
	})

	t.Run("handles string that starts with newline", func(t *testing.T) {
		input := "\nhttp://only.com"
		first, rest := util.SplitFirstLine(input)

		assert.Equal(t, "", first)
		assert.Equal(t, "http://only.com", rest)
	})
}
