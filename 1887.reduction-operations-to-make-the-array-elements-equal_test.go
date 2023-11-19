package main

import "testing"

func TestReductionOperations(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{5, 1, 3}, 3},
		{[]int{1, 1, 1}, 0},
		{[]int{1, 1, 2, 2, 3}, 4},
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5}, 16},
	}
	for _, tt := range tests {
		got := reductionOperations(tt.nums)
		if got != tt.want {
			t.Errorf("reductionOperations(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
