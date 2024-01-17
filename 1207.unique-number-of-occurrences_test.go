package main

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		arr      []int
		expected bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}
	for _, test := range tests {
		actual := uniqueOccurrences(test.arr)
		if actual != test.expected {
			t.Errorf("uniqueOccurrences(%v) expected %v, actual %v", test.arr, test.expected, actual)
		}
	}
}
