package main

import "testing"

func TestAverageofSubtree(t *testing.T) {
	treeOne := &TreeNode2265{Val: 4}
	treeOne.Left = &TreeNode2265{Val: 8}
	treeOne.Right = &TreeNode2265{Val: 5}
	treeOne.Left.Left = &TreeNode2265{Val: 0}
	treeOne.Left.Right = &TreeNode2265{Val: 1}
	treeOne.Right.Right = &TreeNode2265{Val: 6}

	treeTwo := &TreeNode2265{Val: 1}

	if count := averageOfSubtree(treeOne); count != 5 {
		t.Errorf("Expected 6, but got %d", count)
	}

	if count := averageOfSubtree(treeTwo); count != 1 {
		t.Errorf("Expected 1, but got %d", count)
	}
}
