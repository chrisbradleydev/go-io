package writer

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestWriteString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		mockWriter  io.Writer
		expectedN   int
		expectedErr error
	}{
		{
			name:        "Success case",
			input:       "Hello, World!",
			mockWriter:  &bytes.Buffer{},
			expectedN:   13,
			expectedErr: nil,
		},
		{
			name:        "Empty string",
			input:       "",
			mockWriter:  &bytes.Buffer{},
			expectedN:   0,
			expectedErr: nil,
		},
		{
			name:        "Error case",
			input:       "This will fail",
			mockWriter:  &mockErrorWriter{},
			expectedN:   0,
			expectedErr: fmt.Errorf("error occured while writing: %w", errors.New("mock error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, err := WriteString(tt.input, tt.mockWriter)

			if n != tt.expectedN {
				t.Errorf("Expected n to be %d, but got %d", tt.expectedN, n)
			}

			if (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr == nil) {
				t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
			}

			if err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Expected error message '%v', but got '%v'", tt.expectedErr, err)
			}

			if buf, ok := tt.mockWriter.(*bytes.Buffer); ok {
				if buf.String() != tt.input {
					t.Errorf("Expected writer to contain '%s', but got '%s'", tt.input, buf.String())
				}
			}
		})
	}
}

// mockErrorWriter is a mock io.Writer that always returns an error
type mockErrorWriter struct{}

func (m *mockErrorWriter) Write(p []byte) (int, error) {
	return 0, errors.New("mock error")
}
