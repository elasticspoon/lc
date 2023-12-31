package main

import "testing"

func TestTree2Str(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		tree := &TreeNode606{Val: 1, Left: &TreeNode606{Val: 2, Left: &TreeNode606{Val: 4}}, Right: &TreeNode606{Val: 3}}
		got := tree2str(tree)
		want := "1(2(4))(3)"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		tree := &TreeNode606{Val: 1, Left: &TreeNode606{Val: 2, Right: &TreeNode606{Val: 4}}, Right: &TreeNode606{Val: 3}}
		got := tree2str(tree)
		want := "1(2()(4))(3)"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("test three", func(t *testing.T) {
		tree := &TreeNode606{Val: 1, Left: &TreeNode606{Val: 2, Left: &TreeNode606{Val: 3}, Right: &TreeNode606{Val: 4}}}
		got := tree2str(tree)
		want := "1(2(3)(4))"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
