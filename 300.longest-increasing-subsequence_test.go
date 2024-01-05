package main

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}

	for _, tt := range tests {
		got := lengthOfLIS(tt.nums)
		if got != tt.want {
			t.Errorf("lengthOfLIS(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
