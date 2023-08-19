package config

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileOperator struct {
	mock.Mock
}

func (m *MockFileOperator) Open(filename string) (io.ReadCloser, error) {
	args := m.Called(filename)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func (m *MockFileOperator) Write(filename string, data []byte, perm os.FileMode) error {
	args := m.Called(filename, data, perm)
	return args.Error(0)
}

func TestSetAPIKeyWithOperator(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		writeErr error
		wantErr  bool
	}{
		{
			name:    "successful write",
			key:     "sample_api_key",
			wantErr: false,
		},
		{
			name:     "write error",
			key:      "sample_api_key",
			writeErr: errors.New("write error"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOp := new(MockFileOperator)
			mockOp.On("Write", mock.Anything, []byte("api_key="+tt.key), mock.Anything).Return(tt.writeErr)

			err := SetAPIKeyWithOperator(tt.key, mockOp)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockOp.AssertExpectations(t)
		})
	}
}

func TestGetAPIKeyWithOperator(t *testing.T) {
	tests := []struct {
		name    string
		content string
		openErr error
		want    string
		wantErr bool
	}{
		{
			name:    "valid api key",
			content: "api_key=sample_api_key",
			want:    "sample_api_key",
		},
		{
			name:    "missing api_key",
			content: "some_other_key=value",
			wantErr: true,
		},
		{
			name:    "open error",
			openErr: errors.New("open error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOp := new(MockFileOperator)
			rc := io.NopCloser(strings.NewReader(tt.content))
			mockOp.On("Open", mock.Anything).Return(rc, tt.openErr)

			result, err := GetAPIKeyWithOperator(mockOp)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}

			mockOp.AssertExpectations(t)
		})
	}
}
