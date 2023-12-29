package main

import "testing"

func TestMinDifficulty(t *testing.T) {
	tests := []struct {
		jobDifficulty []int
		d             int
		want          int
	}{
		{[]int{1, 2, 3}, 1, 3},
		{[]int{6, 5, 4, 3, 2, 1}, 1, 6},
		{[]int{6, 5, 4, 3, 2, 1}, 2, 7},
		{[]int{6, 5, 4, 3, 2, 1}, 3, 9},
		{[]int{9, 9, 9}, 4, -1},
		{[]int{1, 1, 1}, 3, 3},
		{[]int{7, 1, 7, 1, 7, 1}, 3, 15},
		{[]int{11, 111, 22, 222, 33}, 5, 399},
		{[]int{11, 111, 22, 222, 33, 333, 44, 444}, 6, 843},
	}

	for _, tt := range tests {
		got := minDifficulty(tt.jobDifficulty, tt.d)
		if got != tt.want {
			t.Errorf("minDifficulty(%v, %d) == %d, want %d", tt.jobDifficulty, tt.d, got, tt.want)
		}
	}
}
