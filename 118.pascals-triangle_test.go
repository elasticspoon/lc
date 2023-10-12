// 118.pascals-triangle_test.go

package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		numRows  int
		expected [][]int
	}{
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("numRows=%d", test.numRows), func(t *testing.T) {
			result := generate(test.numRows)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For numRows=%d, expected %v, but got %v", test.numRows, test.expected, result)
			}
		})
	}
}

