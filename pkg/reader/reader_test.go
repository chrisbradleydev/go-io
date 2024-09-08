package reader

import (
	"io"
	"maps"
	"strings"
	"testing"
)

func TestCountLetters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
		wantErr  bool
	}{
		{
			name:     "Basic alphabet count",
			input:    "abcABC",
			expected: map[string]int{"a": 1, "b": 1, "c": 1, "A": 1, "B": 1, "C": 1},
			wantErr:  false,
		},
		{
			name:     "With non-alphabetic characters",
			input:    "Hello, World! 123",
			expected: map[string]int{"H": 1, "e": 1, "l": 3, "o": 2, "W": 1, "r": 1, "d": 1},
			wantErr:  false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]int{},
			wantErr:  false,
		},
		{
			name:  "Long string",
			input: strings.Repeat("abcdefghijklmnopqrstuvwxyz", 100),
			expected: map[string]int{
				"a": 100, "b": 100, "c": 100, "d": 100, "e": 100, "f": 100, "g": 100,
				"h": 100, "i": 100, "j": 100, "k": 100, "l": 100, "m": 100, "n": 100,
				"o": 100, "p": 100, "q": 100, "r": 100, "s": 100, "t": 100, "u": 100,
				"v": 100, "w": 100, "x": 100, "y": 100, "z": 100,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			result, err := CountLetters(reader)

			if (err != nil) != tt.wantErr {
				t.Errorf("CountLetters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !maps.Equal(result, tt.expected) {
				t.Errorf("CountLetters() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCountLettersWithErrorReader(t *testing.T) {
	errReader := &ErrorReader{Err: io.ErrClosedPipe}
	_, err := CountLetters(errReader)
	if err != io.ErrClosedPipe {
		t.Errorf("CountLetters() error = %v, want %v", err, io.ErrClosedPipe)
	}
}

// ErrorReader is a custom io.Reader that always returns an error
type ErrorReader struct {
	Err error
}

func (er *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, er.Err
}
