package main

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tts := []struct {
		Input    string
		Expected time.Time
	}{
		{"Mon, 09 Feb 2026 21:46:23 GMT", time.Date(2026, 2, 9, 21, 46, 23, 0, time.UTC)},
	}

	for _, tt := range tts {
		t.Run(tt.Input, func(t *testing.T) {
			t.Parallel()

			actual := parseInput(tt.Input)
			if !actual.Equal(tt.Expected) {
				t.Errorf("failed to parse %q, got %v", tt.Input, actual)
			}
		})
	}
}
