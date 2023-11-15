package main

import "testing"

func TestMaxElAfterDecandRearrange(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{2, 2, 1, 2, 1}, 2},
		{[]int{100, 1, 1000}, 3},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{73, 98, 9, 7, 5, 1, 100, 120}, 8},
		{[]int{73, 98, 9}, 3},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1},
	}

	for _, tt := range tests {
		if got := maximumElementAfterDecrementingAndRearranging(tt.arr); got != tt.want {
			t.Errorf("maxElAfterDecandRearrange(%v) = %v, want %v", tt.arr, got, tt.want)
		}
	}
}
