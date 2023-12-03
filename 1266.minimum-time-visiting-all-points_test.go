package main

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7},
		{[][]int{{3, 2}, {-2, 2}}, 5},
	}
	for _, tt := range tests {
		got := minTimeToVisitAllPoints(tt.points)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
