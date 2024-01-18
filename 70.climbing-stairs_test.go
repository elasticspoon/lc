package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		height   int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{45, 1836311903},
	}

	for _, test := range tests {
		actual := climbStairs(test.height)
		if actual != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, actual)
		}
	}
}
