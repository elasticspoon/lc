package main

import "testing"

func TestMaxScore1793(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 5},
		{[]int{1, 4, 3, 7, 4, 5}, 3, 15},
		{[]int{5, 5, 4, 5, 4, 1, 1, 1}, 0, 20},
		{[]int{6569, 9667, 3148, 7698, 1622, 2194, 793, 9041, 1670, 1872}, 5, 9732},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 30},
		{[]int{8182, 1273, 9847, 6230, 52, 1467, 6062, 726, 4852, 4507, 2460, 2041, 500, 1025, 5524}, 8, 9014},
	}

	for _, tt := range tests {
		got := maximumScore(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("maximumScore(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
