package main

import "testing"

func TestNumBusesToDest(t *testing.T) {
	tests := []struct {
		routes [][]int
		source int
		target int
		want   int
	}{
		{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6, 2},
		{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 5, -1},
		{[][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12, -1},
	}
	for _, tt := range tests {
		got := numBusesToDestination(tt.routes, tt.source, tt.target)
		if got != tt.want {
			t.Errorf("numBusesToDestination(%v, %v, %v) = %v, want %v", tt.routes, tt.source, tt.target, got, tt.want)
		}
	}
}
