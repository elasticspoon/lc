package main

import "testing"

func TestMaxScore(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"011101", 5},
		{"00111", 5},
		{"1111", 3},
		{"00", 1},
	}
	for i, tt := range tests {
		got := maxScore(tt.s)
		if got != tt.want {
			t.Errorf("%d: maxScore(%v) = %d, want %d", i, tt.s, got, tt.want)
		}
	}
}
