package main

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		corridor string
		want     int
	}{
		{"SSPPSPS", 3},
		{"PPSPSP", 1},
		{"S", 0},
		{"P", 0},
		{"SSPPSSPPSS", 9},
		{"SSPPSS", 3},
	}

	for _, tt := range tests {
		got := numberOfWays(tt.corridor)
		if got != tt.want {
			t.Errorf("numberOfWays(%v) = %v, want %v", tt.corridor, got, tt.want)
		}
	}
}
