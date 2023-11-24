package main

import "testing"

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		piles []int
		want  int
	}{
		{[]int{2, 4, 1, 2, 7, 8}, 9},
		{[]int{2, 4, 5}, 4},
		{[]int{9, 8, 7, 6, 5, 1, 2, 3, 4}, 18},
	}
	for _, tt := range tests {
		got := maxCoins(tt.piles)
		if got != tt.want {
			t.Errorf("maxCoins(%v) = %v, want %v", tt.piles, got, tt.want)
		}
	}
}
