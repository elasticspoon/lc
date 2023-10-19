package main

import "testing"

func TestValidateBinaryTreeNodes(t *testing.T) {
	tests := []struct {
		n          int
		leftChild  []int
		rightChild []int
		expected   bool
	}{
		{4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}, true},
		{4, []int{-1, -1, 0, 1}, []int{-1, -1, -1, 2}, true}, // root node is not first node
		{4, []int{-1, -1, 0, 1}, []int{-1, -1, 1, 2}, false}, // root node is not first node, but has cycle
		{4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}, false},
		{2, []int{1, 0}, []int{-1, -1}, false}, // finds a cycle
		{6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1}, false},
		{4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}, false},
		{5, []int{4, -1, 3, -1, -1}, []int{1, -1, -1, -1, -1}, false},
		{5, []int{3, -1, 1, -1, -1}, []int{2, -1, 4, -1, -1}, true},
	}

	for _, test := range tests {
		actual := validateBinaryTreeNodes(test.n, test.leftChild, test.rightChild)
		if actual != test.expected {
			t.Errorf("validateBinaryTreeNodes(%d, %v, %v) expected %v, actual %v", test.n, test.leftChild, test.rightChild, test.expected, actual)
		}
	}
}
