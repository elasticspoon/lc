package main

import (
	"reflect"
	"testing"
)

func TestSortByBits(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, expected: []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
		{input: []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}, expected: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}},
	}

	for _, test := range tests {
		got := sortByBits(test.input)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("sortByBits(%v) = %v; want %v", test.input, got, test.expected)
		}

	}
}
