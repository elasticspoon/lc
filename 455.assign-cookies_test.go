package main

import "testing"

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		g    []int
		s    []int
		want int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}, 2},
		{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8, 9, 10}, 4},
	}

	for _, tt := range tests {
		got := findContentChildren(tt.g, tt.s)
		if got != tt.want {
			t.Errorf("findContentChildren(%v, %v) = %v, want %v", tt.g, tt.s, got, tt.want)
		}
	}
}
