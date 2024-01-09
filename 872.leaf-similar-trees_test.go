package main

import (
	"testing"
)

func TestLeafSimilar(t *testing.T) {
	tests := []struct {
		tree1 []any
		tree2 []any
		want  bool
	}{
		{[]any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}, []any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}, true},
		{[]any{1}, []any{1}, true},
		{[]any{1, 3, 2}, []any{1, 2, 3}, false},
	}

	for _, tt := range tests {
		tree1 := buildTree(tt.tree1, 0)
		tree2 := buildTree(tt.tree2, 0)
		got := leafSimilar(tree1, tree2)
		if got != tt.want {
			t.Errorf("leafSimilar(%v, %v) = %v, want %v", tree1, tree2, got, tt.want)
		}
	}
}

func buildTree(vals []any, index int) *TreeNode {
	if index >= len(vals) || vals[index] == nil {
		return nil
	}

	v := vals[index].(int)
	root := &TreeNode{Val: v}
	root.Left = buildTree(vals, 2*index+1)
	root.Right = buildTree(vals, 2*index+2)
	return root
}
