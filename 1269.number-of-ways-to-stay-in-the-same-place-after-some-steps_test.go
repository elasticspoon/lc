package main

import (
	"testing"
)

func TestNumWays(t *testing.T) {
	tests := []struct {
		steps    int
		arrLen   int
		expected int
	}{
		{3, 2, 4},
		{3, 1, 1},
		{2, 4, 2},
		{4, 2, 8},
		{4, 3, 9},
	}

	for _, test := range tests {
		actual := numWays(test.steps, test.arrLen)
		if actual != test.expected {
			t.Errorf("numWays(%d, %d) expected %d, actual %d", test.steps, test.arrLen, test.expected, actual)
		}
	}
}
