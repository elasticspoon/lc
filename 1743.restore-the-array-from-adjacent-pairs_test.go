package main

import (
	"reflect"
	"testing"
)

func TestRestoreArray(t *testing.T) {
	tests := []struct {
		adjacentPairs [][]int
		want          []int
	}{
		{[][]int{{2, 1}, {3, 4}, {3, 2}}, []int{1, 2, 3, 4}},
		{[][]int{{4, -2}, {1, 4}, {-3, 1}}, []int{-2, 4, 1, -3}},
		{[][]int{{100000, -100000}}, []int{100000, -100000}},
		{[][]int{{4, -10}, {-1, 3}, {4, -3}, {-3, 3}}, []int{-1, 3, -3, 4, -10}},
	}

	for _, tt := range tests {
		got := restoreArray(tt.adjacentPairs)
		if len(got) != len(tt.want) {
			t.Errorf("len wrong: restoreArray(%v) = %v, want %v", tt.adjacentPairs, got, tt.want)
		}
		if !reflect.DeepEqual(got, tt.want) {
			for i, v := range got {
				if v != tt.want[len(tt.want)-1-i] {
					t.Errorf("restoreArray(%v) = %v, want %v", tt.adjacentPairs, got, tt.want)
					break
				}
			}
		}
	}
}
