package util_test

import (
	"reflect"
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

type TestStruct struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

func TestUnmarshalJSONBlock(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected TestStruct
		wantErr  bool
	}{
		{
			name:  "Valid JSON with ```json marker",
			input: "```json\n{\n  \"title\": \"Test Title\",\n  \"summary\": \"Test summary text.\"\n}\n```",
			expected: TestStruct{
				Title:   "Test Title",
				Summary: "Test summary text.",
			},
			wantErr: false,
		},
		{
			name:  "Valid JSON with ``` only",
			input: "```\n{\n  \"title\": \"Title Only\",\n  \"summary\": \"Just summary.\"\n}\n```",
			expected: TestStruct{
				Title:   "Title Only",
				Summary: "Just summary.",
			},
			wantErr: false,
		},
		{
			name:  "Valid JSON without any markers",
			input: "{ \"title\": \"No Markers\", \"summary\": \"Still works.\" }",
			expected: TestStruct{
				Title:   "No Markers",
				Summary: "Still works.",
			},
			wantErr: false,
		},
		{
			name:    "Empty JSON block",
			input:   "```json\n\n```",
			wantErr: true,
		},
		{
			name:    "Invalid JSON block",
			input:   "```json\ninvalid json\n```",
			wantErr: true,
		},
		{
			name:    "Completely empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:  "Extra whitespace around markers",
			input: "   ```json\n  {\n\"title\": \"Trimmed\", \"summary\": \"Works fine\"  }\n```   ",
			expected: TestStruct{
				Title:   "Trimmed",
				Summary: "Works fine",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := util.UnmarshalJSONBlock[TestStruct](tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSONBlock() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("UnmarshalJSONBlock() = %+v, expected %+v", got, tt.expected)
			}
		})
	}
}
