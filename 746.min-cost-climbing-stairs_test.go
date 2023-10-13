package main

import "testing"

func TestMinClimbCost(t *testing.T) {
	tests := []struct {
		cost     []int
		expected int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{[]int{0, 0, 0, 1}, 0},
		{[]int{0, 0, 1, 0}, 0},
		{[]int{0, 1, 0, 0}, 0},
		{[]int{1, 0, 0, 0}, 0},
		{[]int{10, 15, 20, 25, 30}, 40},
	}

	for _, tt := range tests {
		actual := minCostClimbingStairs(tt.cost)
		if actual != tt.expected {
			t.Errorf("minCostClimbingStairs(%v): expected %d, actual %d", tt.cost, tt.expected, actual)
		}
	}
}
