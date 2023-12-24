package main

import "testing"

func TestMinOperations(t *testing.T) {
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
		got := minOperations(tt.s)
		if got != tt.want {
			t.Errorf("minOperations(%v)=%v, want %v", tt.s, got, tt.want)
		}
	}
}
