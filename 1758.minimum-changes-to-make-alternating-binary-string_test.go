package main

import "testing"

func TestMinOperations1758(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"0100", 1},
		{"10", 0},
		{"1111", 2},
		{"10010100", 3},
	}

	for _, tt := range tests {
		got := minOperations1758(tt.s)
		if got != tt.want {
			t.Errorf("minOperations(%v)=%v, want %v", tt.s, got, tt.want)
		}
	}
}
