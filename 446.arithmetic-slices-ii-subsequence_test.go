package main

import "testing"

func TestNumberofArithmeticSlices(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// {[]int{2, 4, 6, 8, 10}, 7},
		// {[]int{-4, -2, 0, 2, 4}, 7},
		// {[]int{4, 2, 0, -2, -4}, 7},
		// {[]int{7, 7, 7, 7, 7}, 16},
		// {[]int{0, 1, 2, 2, 2}, 4},
		// {[]int{2, 2, 3, 4}, 2},
		// {[]int{0, 1, 2, 2, 2, 3, 4}, 6},
	}

	for _, tt := range tests {
		got := numberOfArithmeticSlices(tt.nums)
		if got != tt.want {
			t.Errorf("numberOfArithmeticSlices(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func TestZeroDiffSubseqs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		// {3, 1},
		// {4, 5},
		// {5, 16},
	}

	for _, tt := range tests {
		got := zeroDiffSubseqs(tt.n)
		if got != tt.want {
			t.Errorf("zeroDiffSubseqs(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
