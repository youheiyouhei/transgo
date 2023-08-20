package cmd

import (
	"testing"

	"github.com/youheiyouhei/transgo/interfaces"
)

func TestFormatSupportedLanguages(t *testing.T) {
	tests := []struct {
		name     string
		input    interfaces.SupportedLanguages
		expected string
	}{
		{
			name: "multiple languages",
			input: interfaces.SupportedLanguages{
				{Name: "English", Code: "EN"},
				{Name: "Japanese", Code: "JP"},
			},
			expected: "Available languages:\n- English (EN)\n- Japanese (JP)\n",
		},
		{
			name:     "no languages",
			input:    interfaces.SupportedLanguages{},
			expected: "Available languages:\n",
		},
		{
			name: "single language",
			input: interfaces.SupportedLanguages{
				{Name: "English", Code: "EN"},
			},
			expected: "Available languages:\n- English (EN)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := formatSupportedLanguages(tt.input)
			if output != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, output)
			}
		})
	}
}
