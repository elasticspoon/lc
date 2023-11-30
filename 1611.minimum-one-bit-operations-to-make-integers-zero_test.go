package main

import "testing"

func TestMinimumOneBitOperations(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{3, 2},
		{6, 4},
		{9, 14},
		{14, 11},
		{46, 52},
		{19, 29},
		{22, 27},
		{333, 393},
	}
	for _, tt := range tests {
		got := minimumOneBitOperations(tt.n)
		if got != tt.want {
			t.Errorf("minimumOneBitOperations(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
