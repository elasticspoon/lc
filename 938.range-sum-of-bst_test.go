package main

import "testing"

func TestRangeSumBST(t *testing.T) {
	tree := &TreeNode938{
		10,
		&TreeNode938{5, &TreeNode938{3, nil, nil}, &TreeNode938{7, nil, nil}},
		&TreeNode938{15, nil, &TreeNode938{18, nil, nil}},
	}

	tests := []struct {
		start int
		end   int
		want  int
	}{
		{7, 15, 32},
		{6, 10, 17},
		{5, 15, 37},
		{3, 18, 58},
	}

	for _, tt := range tests {
		got := rangeSumBST(tree, tt.start, tt.end)
		if got != tt.want {
			t.Errorf("rangeSumBST(%v, %v, %v) = %v, want %v", tree, tt.start, tt.end, got, tt.want)
		}
	}
}
