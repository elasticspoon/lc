package main

import (
	"testing"
)

func TestFindMatrix(t *testing.T) {
	unorderedEql := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		mapA := make(map[int]bool)
		for _, v := range a {
			mapA[v] = true
		}
		for _, v := range b {
			if _, ok := mapA[v]; !ok {
				return false
			}
		}
		return true
	}

	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}}},
		{[]int{1, 3, 4, 1, 2, 3, 1}, [][]int{{4, 2, 1, 3}, {1, 3}, {1}}},
	}

	for _, tt := range tests {
		got := findMatrix(tt.nums)

		for i := 0; i < len(got); i++ {
			if !unorderedEql(got[i], tt.want[i]) {
				t.Errorf("findMatrix(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		}
	}
}
