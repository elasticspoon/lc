package main

import "testing"

func TestFindSpecialInteger(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 2, 6, 6, 6, 6, 7, 10}, 6},
		{[]int{1, 2, 2, 6, 6, 6, 7, 10}, 6},
		{[]int{1, 2, 6, 6, 7, 10}, 6},
		{[]int{1, 1}, 1},
	}

	for _, tt := range tests {
		got := findSpecialInteger(tt.arr)
		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
