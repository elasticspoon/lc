package main

import "testing"

func TestIsReachableAtTime(t *testing.T) {
	tests := []struct {
		sx   int
		sy   int
		fx   int
		fy   int
		t    int
		want bool
	}{
		{2, 4, 7, 7, 6, true},
		{3, 1, 7, 3, 3, false},
		{0, 0, 1, 1, 1, true},
		{2, 2, 4, 4, 2, true},
		{1, 2, 1, 2, 1, false},
	}

	for _, tt := range tests {
		got := isReachableAtTime(tt.sx, tt.sy, tt.fx, tt.fy, tt.t)
		if got != tt.want {
			t.Errorf("isReachableAtTime(%v, %v, %v, %v, %v) = %v, want %v", tt.sx, tt.sy, tt.fx, tt.fy, tt.t, got, tt.want)
		}
	}
}
