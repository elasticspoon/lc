package main

import "testing"

func TestNumberofMatches(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{7, 6},
		{14, 13},
		{1, 0},
		{200, 199},
	}

	for _, tt := range tests {
		got := numberOfMatches(tt.in)
		if got != tt.want {
			t.Errorf("numberofMatches(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
