package main

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 2}, 12},
		{[]int{1, 5, 4, 5}, 16},
		{[]int{3, 7}, 12},
	}

	for _, tt := range tests {
		got := maxProduct(tt.nums)
		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
