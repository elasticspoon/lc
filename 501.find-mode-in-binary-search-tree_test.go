package main

import (
	"reflect"
	"testing"
)

func TestFindMode(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 2}, []int{2}},
		{[]int{1, 1, 2, 2}, []int{1, 2}},
		{[]int{1, 2, 2, 3, 3, 3}, []int{3}},
		{[]int{0}, []int{0}},
	}

	for _, tt := range tests {
		root := buildTree501(tt.input)
		got := findMode(root)
		if reflect.DeepEqual(got, tt.want) {
			t.Errorf("findMode(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func buildTree501(vals []int) *TreeNode501 {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode501{Val: vals[0]}
	for i := 1; i < len(vals); i++ {
		root.Insert(vals[i])
	}
	return root
}
