package main

import "testing"

func TestCountHomogenous(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"y", 1},
		{"abbcccaa", 13},
		{"xy", 2},
		{"zzzzz", 15},
		{"zzzzzzz", 28},
		{"zzzzzzzz", 36},
	}

	for _, tt := range tests {
		got := countHomogenous(tt.s)
		if got != tt.want {
			t.Errorf("countHomogenous(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
