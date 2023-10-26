package main

import "testing"

func TestNumFactoredBinaryTrees(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{arr: []int{2, 4}, want: 3},
		{arr: []int{2, 4, 5, 10}, want: 7},
		{arr: []int{2, 4, 5, 10, 25}, want: 9},
		{arr: []int{2, 4, 5, 10, 25, 8, 16, 32}, want: 80},
		{arr: []int{2, 5, 10, 25}, want: 7},
		{arr: []int{2, 4, 8, 16, 32}, want: 74},
	}

	for _, tt := range tests {
		result := numFactoredBinaryTrees(tt.arr)
		if result != tt.want {
			t.Errorf("numFactoredBinaryTrees(%v) = %v, want %v", tt.arr, result, tt.want)
		}
	}
}
