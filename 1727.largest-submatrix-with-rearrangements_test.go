package main

import "testing"

func TestLargestSubmatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{1}, {0}, {1}}, 1},
		{[][]int{{1, 0, 1, 0, 1}}, 3},
		{[][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}, 4},
		{[][]int{{1, 1, 0}, {1, 0, 1}}, 2},
		{[][]int{{0, 0}, {0, 0}}, 0},
	}

	for _, tt := range tests {
		got := largestSubmatrix(tt.matrix)
		if got != tt.want {
			t.Errorf("largestSubmatrix(%v) = %v, want %v", tt.matrix, got, tt.want)
		}
	}
}
