package main

import (
	"testing"
)

func TestRecDist(t *testing.T) {
	tests := []struct {
		grid  [][]int
		point []int
		want  int
	}{
		{[][]int{{0, 2, 1}, {1, 1, 1}, {1, 1, 1}}, []int{0, 0}, 1},
		{[][]int{{1, 1, 1}, {1, 1, 1}, {0, 1, 2}}, []int{2, 0}, 2},
		{[][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 2}}, []int{0, 0}, 4},
		{[][]int{{3, 2, 0}, {0, 1, 0}, {0, 3, 0}}, []int{0, 2}, 1},
	}

	for i, tt := range tests {
		visited := []int{}
		got, _ := recDist(tt.grid, tt.point, visited)
		if got != tt.want {
			t.Errorf("%d: recDist(%v, %v) = %d, want %d", i, tt.grid, tt.point, got, tt.want)
		}
	}
}

func TestMinimumMoves(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		// {[][]int{{0, 2, 1}, {1, 1, 1}, {1, 1, 1}}, 1},
		// {[][]int{{1, 1, 0}, {1, 1, 1}, {1, 2, 1}}, 3},
		// {[][]int{{1, 3, 0}, {1, 0, 0}, {1, 0, 3}}, 4},
		// {[][]int{{3, 2, 0}, {0, 1, 0}, {0, 3, 0}}, 7},
	}

	for i, tt := range tests {
		got := minimumMoves(tt.grid)
		if got != tt.want {
			t.Errorf("%d: minimumMoves(%v) = %d, want %d", i, tt.grid, got, tt.want)
		}
	}
}
