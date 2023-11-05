package main

import "testing"

func TestGetLastMoment(t *testing.T) {
	tests := []struct {
		n        int
		left     []int
		right    []int
		expected int
	}{
		{4, []int{4, 3}, []int{0, 1}, 4},
		{7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}, 7},
		{7, []int{0, 1, 2, 3, 4, 5, 6, 7}, []int{}, 7},
		{9, []int{5}, []int{4}, 5},
		{9, []int{5}, []int{4, 8}, 5},
	}

	for _, tt := range tests {
		actual := getLastMoment(tt.n, tt.left, tt.right)
		if actual != tt.expected {
			t.Errorf("getLastMoment(%d, %v, %v): expected %d, actual %d",
				tt.n, tt.left, tt.right, tt.expected, actual)
		}
	}
}
