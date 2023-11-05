package main

import "testing"

func TestGetWinner(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{2, 1, 3, 5, 4, 6, 7}, 2, 5},
		{[]int{3, 2, 1}, 10, 3},
		{[]int{1, 9, 8, 2, 3, 7, 6, 4, 5}, 7, 9},
		{[]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000, 99},
	}

	for _, tt := range tests {
		actual := getWinner(tt.arr, tt.k)
		if actual != tt.expected {
			t.Errorf("getWinner(%v, %d): expected %d, actual %d",
				tt.arr, tt.k, tt.expected, actual)
		}
	}
}
