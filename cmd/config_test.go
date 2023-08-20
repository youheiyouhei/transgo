package cmd

import (
	"testing"
)

func TestParseKeyValue(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expectedK  string
		expectedV  string
		shouldFail bool
	}{
		{
			name:       "valid format",
			input:      "api_key=12345",
			expectedK:  "api_key",
			expectedV:  "12345",
			shouldFail: false,
		},
		{
			name:       "invalid format no value",
			input:      "api_key=",
			expectedK:  "",
			expectedV:  "",
			shouldFail: true,
		},
		{
			name:       "invalid format no key",
			input:      "=12345",
			expectedK:  "",
			expectedV:  "",
			shouldFail: true,
		},
		{
			name:       "invalid format no equals sign",
			input:      "apikey12345",
			expectedK:  "",
			expectedV:  "",
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			k, v, err := parseKeyValue(tt.input)
			if tt.shouldFail {
				if err == nil {
					t.Errorf("Expected an error but did not get one")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got %v", err)
				}
				if k != tt.expectedK || v != tt.expectedV {
					t.Errorf("Expected key=%s, value=%s but got key=%s, value=%s", tt.expectedK, tt.expectedV, k, v)
				}
			}
		})
	}
}
