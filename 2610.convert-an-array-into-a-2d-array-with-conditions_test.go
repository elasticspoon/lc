package main

import (
	"reflect"
	"testing"
)

func TestFindMatrix(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}}},
		{[]int{1, 3, 4, 1, 2, 3, 1}, [][]int{{1, 3, 4, 2}, {1, 3}, {1}}},
	}

	for _, tt := range tests {
		got := findMatrix(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findMatrix(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
