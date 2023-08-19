package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTranslator is our mock object implementing the Translator interface
type MockTranslator struct {
	mock.Mock
}

func (m *MockTranslator) Translate(texts []string, source, target string) (string, error) {
	args := m.Called(texts, source, target)
	return args.String(0), args.Error(1)
}

func TestHandleTranslation(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		source     string
		target     string
		mockReturn string
		mockErr    error
		expected   string
		err        error
	}{
		{
			name:       "successful translation",
			text:       "Hello, world!",
			source:     "en",
			target:     "ja",
			mockReturn: "Mocked translation",
			mockErr:    nil,
			expected:   "Mocked translation",
			err:        nil,
		},
		{
			name:     "failed translation",
			text:     "Hello, world!",
			source:   "en",
			target:   "ja",
			mockErr:  errors.New("mocked error"),
			expected: "",
			err:      errors.New("mocked error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the mock to return specific values
			client := new(MockTranslator)
			client.On("Translate", []string{tt.text}, tt.source, tt.target).Return(tt.mockReturn, tt.mockErr)

			got, err := handleTranslation(tt.text, tt.source, tt.target, client)

			assert.Equal(t, tt.expected, got)
			assert.Equal(t, tt.err, err)

			// Verify that the Translate function was called with the expected arguments
			client.AssertExpectations(t)
		})
	}
}
