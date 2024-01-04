package main

import "testing"

func TestMinOperations2870(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 3, 2, 2, 4, 2, 3, 4}, 4},
		{[]int{2, 1, 2, 2, 3, 3}, -1},
		{[]int{1}, -1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 1}, 1},
		{[]int{1, 1, 1, 1}, 2},
		{[]int{1, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 3},
	}

	for _, tt := range tests {
		got := minOperations(tt.nums)
		if got != tt.want {
			t.Errorf("minOperations(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
