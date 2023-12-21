package main

import "testing"

func TestMaxWidthOfVerticalArea(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}, 1},
		{[][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}, 3},
	}
	for i, tt := range tests {
		got := maxWidthOfVerticalArea(tt.points)
		if got != tt.want {
			t.Errorf("%d: maxWidthOfVerticalArea(%v) = %d, want %d", i, tt.points, got, tt.want)
		}
	}
}
