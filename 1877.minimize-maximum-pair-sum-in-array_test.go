package main

import "testing"

func TestMinPairSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 5, 2, 3}, 7},
		{[]int{3, 5, 4, 2, 4, 6}, 8},
	}
	for _, tt := range tests {
		got := minPairSum(tt.nums)
		if got != tt.want {
			t.Errorf("minPairSum(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
