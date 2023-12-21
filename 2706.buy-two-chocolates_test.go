package main

import "testing"

func TestBuyChoco(t *testing.T) {
	tests := []struct {
		prices []int
		money  int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 2},
		{[]int{1, 2, 3, 4, 5}, 4, 1},
		{[]int{1, 2, 3, 4, 5}, 3, 0},
		{[]int{1, 2, 2}, 3, 0},
		{[]int{3, 2, 3}, 3, 3},
	}

	for i, tt := range tests {
		got := buyChoco(tt.prices, tt.money)
		if got != tt.want {
			t.Errorf("%d: buyChoco(%v, %d) = %d, want %d", i, tt.prices, tt.money, got, tt.want)
		}
	}
}
