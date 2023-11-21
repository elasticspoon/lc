package main

import "testing"

func TestCountNicePairs(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{42, 11, 1, 97}, 2},
		{[]int{13, 10, 35, 24, 76}, 4},
		{[]int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6}, 55},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 66},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 78},
	}
	longTest := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		longTest[i] = i
	}
	tests = append(tests, struct {
		nums []int
		want int
	}{longTest, 19468545})

	for _, tt := range tests {
		got := countNicePairs(tt.nums)
		if got != tt.want {
			t.Errorf("countNicePairs(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func TestReverseInt(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{123, 321},
		{120, 21},
		{0, 0},
		{1, 1},
		{10, 1},
		{9812374, 4732189},
	}

	for _, tt := range tests {
		got := reverseInt(tt.n)
		if got != tt.want {
			t.Errorf("reverseInt(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
