package main

import "testing"

func TestConstrainedSubsetSum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 2, -10, 5, 20}, 2, 37},
		{[]int{-1, -2, -3}, 1, -1},
		{[]int{10, -2, -10, -5, 20}, 2, 23},
		{[]int{10, -2, -10, -5, 20}, 3, 28},
		{[]int{10, -2, -10, -5, 20}, 4, 30},
	}

	for _, tt := range tests {
		got := constrainedSubsetSum(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("constrainedSubsetSum(%v, %v) = %v; want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
