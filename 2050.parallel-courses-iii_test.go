package main

import (
	"testing"
)

func TestMinimumTime(t *testing.T) {
	tests := []struct {
		n         int
		relations [][]int
		time      []int
		expected  int
	}{
		{3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 1}, 4},
		{3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}, 8},
		{5, [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}, []int{1, 2, 3, 4, 5}, 12},
		{3, [][]int{{2, 1}, {1, 3}}, []int{2, 2, 2}, 6},
		{2, [][]int{{2, 1}}, []int{10000, 10000}, 20000},
		{1, [][]int{}, []int{1}, 1},
	}

	for _, test := range tests {
		actual := minimumTime(test.n, test.relations, test.time)
		if actual != test.expected {
			t.Errorf("minimumTime(%d, %v, %v) = %d; expected %d", test.n, test.relations, test.time, actual, test.expected)
		}
	}
}
