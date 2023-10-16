package main

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		rowIndex int
		expected []int
	}{
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
		{0, []int{1}},
		{1, []int{1, 1}},
		{9, []int{1, 9, 36, 84, 126, 126, 84, 36, 9, 1}},
	}

	for _, test := range tests {
		actual := getRow(test.rowIndex)
		if len(actual) != len(test.expected) {
			t.Errorf("getRow(%d) expected %v, actual %v", test.rowIndex, test.expected, actual)
		} else {
			for i := range actual {
				if actual[i] != test.expected[i] {
					t.Errorf("getRow(%d) expected %v, actual %v", test.rowIndex, test.expected, actual)
					break
				}
			}
		}
	}
}
