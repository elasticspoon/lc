package main

import "testing"

func TestKnightDialer(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 10},
		{2, 20},
		{3, 46},
		{4, 104},
		{3131, 136006598},
	}
	for _, tt := range tests {
		got := knightDialer(tt.n)
		if got != tt.want {
			t.Errorf("knightDialer(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
