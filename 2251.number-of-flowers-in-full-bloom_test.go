package main

import (
	"reflect"
	"testing"
)

func TestFullBloomFlowers(t *testing.T) {
	tests := []struct {
		flowers  [][]int
		people   []int
		expected []int
	}{
		{
			flowers:  [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}},
			people:   []int{2, 3, 7, 11},
			expected: []int{1, 2, 2, 2},
		},
		{
			flowers:  [][]int{{3, 7}, {1, 6}, {9, 12}, {4, 13}},
			people:   []int{2, 3, 7, 11},
			expected: []int{1, 2, 2, 2},
		},
		{
			flowers:  [][]int{{3, 7}, {1, 6}, {9, 12}, {4, 13}},
			people:   []int{2, 3, 7, 11},
			expected: []int{1, 2, 2, 2},
		},
		{
			flowers:  [][]int{{1, 10}, {3, 3}},
			people:   []int{3, 3, 2},
			expected: []int{2, 2, 1},
		},
	}

	for _, tt := range tests {
		t.SkipNow()
		actual := fullBloomFlowers(tt.flowers, tt.people)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("expected %v, actual %v", tt.expected, actual)
		}

	}
}
