package main

import "testing"

func TestIsPathCrossing(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"NES", false},
		{"NESWW", true},
	}

	for _, tt := range tests {
		got := isPathCrossing(tt.path)
		if got != tt.want {
			t.Errorf("isPathCrossing(%v)=%v, want %v", tt.path, got, tt.want)
		}
	}
}
