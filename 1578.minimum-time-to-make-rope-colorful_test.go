package main

import "testing"

func TestMinCost(t *testing.T) {
	cases := []struct {
		colors     string
		neededTime []int
		want       int
	}{
		{"abaac", []int{1, 2, 3, 4, 5}, 3},
		{"abc", []int{1, 2, 3}, 0},
		{"aabaa", []int{1, 2, 3, 4, 1}, 2},
		{"aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}, 26},
	}
	for _, c := range cases {
		got := minCost(c.colors, c.neededTime)
		if c.want != got {
			t.Errorf("minCost(%q, %v) == %v, want %v", c.colors, c.neededTime, got, c.want)
		}
	}
}
