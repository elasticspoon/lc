package main

import "testing"

func TestSumSubarrayMins(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{3, 1, 2, 4}, 17},
		{[]int{11, 81, 94, 43, 3}, 444},
		{[]int{71, 55, 82, 55}, 593},
	}

	for _, test := range tests {
		actual := sumSubarrayMins(test.arr)
		if actual != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, actual)
		}
	}
}
