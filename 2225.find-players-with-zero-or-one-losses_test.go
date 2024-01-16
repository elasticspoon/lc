package main

import (
	"reflect"
	"testing"
)

func TestFindWinners(t *testing.T) {
	tests := []struct {
		matches  [][]int
		expected [][]int
	}{
		{[][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}, [][]int{{1, 2, 10}, {4, 5, 7, 8}}},
		{[][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}, [][]int{{1, 2, 5, 6}, {}}},
	}

	for _, test := range tests {
		actual := findWinners(test.matches)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("findWinners(%v) expected %v, actual %v", test.matches, test.expected, actual)
		}
	}
}
