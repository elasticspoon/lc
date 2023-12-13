package main

import "testing"

func TestNumSpecial(t *testing.T) {
	tests := []struct {
		mat  [][]int
		want int
	}{
		{[][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}, 1},
		{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
		{[][]int{{0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 1}, {0, 0, 0, 0, 1, 0, 0, 0}, {1, 0, 0, 0, 1, 0, 0, 0}, {0, 0, 1, 1, 0, 0, 0, 0}}, 1},
	}

	for _, tt := range tests {
		got := numSpecial(tt.mat)
		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
