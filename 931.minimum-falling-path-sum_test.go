package main

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected int
	}{
		{[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13},
		{[][]int{{-19, 57}, {-40, -5}}, -59},
		{[][]int{{-48}}, -48},
		{[][]int{{-73, 61, 43, -48, -36}, {3, 30, 27, 57, 10}, {96, -76, 84, 59, -15}, {5, -49, 76, 31, -7}, {97, 91, 61, -46, 67}}, -134},
	}

	for _, test := range tests {
		actual := minFallingPathSum(test.matrix)
		if actual != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, actual)
		}
	}
}
