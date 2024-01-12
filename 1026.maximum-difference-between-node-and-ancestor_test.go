package main

import "testing"

func TestMaxAncestorDiff(t *testing.T) {
	tests := []struct {
		tree []any
		want int
	}{
		{[]any{8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, 13}, 7},
		{[]any{1, nil, 2, nil, 0, 3}, 3},
	}

	for _, tt := range tests {
		tree := buildTree(tt.tree, 0)
		got := maxAncestorDiff(tree)
		if got != tt.want {
			t.Errorf("maxAncestorDiff(%v) = %v, want %v", tt.tree, got, tt.want)
		}
	}
}
