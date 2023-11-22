package main

import (
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	tests := []struct {
		nums [][]int
		want []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 4, 2, 7, 5, 3, 8, 6, 9}},
		{[][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}, []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}},
	}

	for _, tt := range tests {
		got := findDiagonalOrder(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findDiagonalOrder(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
