package main

import "testing"

func TestEliminateMaximum(t *testing.T) {
	tests := []struct {
		dist  []int
		speed []int
		want  int
	}{
		{[]int{1, 3, 4}, []int{1, 1, 1}, 3},
		{[]int{1, 1, 2, 3}, []int{1, 1, 1, 1}, 1},
		{[]int{3, 2, 4}, []int{5, 3, 2}, 1},
		{[]int{3, 5, 7, 4, 5}, []int{2, 3, 6, 3, 2}, 2},
		{[]int{5, 4, 7, 5, 3}, []int{2, 3, 6, 3, 2}, 2},
		{[]int{4, 2, 8}, []int{2, 1, 4}, 2},
	}

	for _, tt := range tests {
		got := eliminateMaximum(tt.dist, tt.speed)
		if got != tt.want {
			t.Errorf("eliminateMaximum(%v, %v) = %v, want %v", tt.dist, tt.speed, got, tt.want)
		}
	}
}
