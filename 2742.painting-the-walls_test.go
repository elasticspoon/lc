package main

import (
	"testing"
)

func TestPaintWalls(t *testing.T) {
	tests := []struct {
		cost     []int
		time     []int
		expected int
	}{
		{[]int{1, 3, 5, 7, 9}, []int{5, 4, 3, 2, 1}, 1},
		{[]int{1, 3, 5, 7, 9}, []int{1, 2, 3, 4, 5}, 4},
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 2}, 3},
		{[]int{2, 3, 4, 2}, []int{1, 1, 1, 1}, 4},
		{[]int{42, 8, 28, 35, 21, 13, 21, 35}, []int{2, 1, 1, 1, 2, 1, 1, 2}, 63},
	}

	for _, tt := range tests {
		actual := paintWalls(tt.cost, tt.time)
		if actual != tt.expected {
			t.Errorf("paintWalls(cost: %v, time: %v): expected %d, actual %d", tt.cost, tt.time, tt.expected, actual)
		}
	}
}
