package main

import (
	"reflect"
	"testing"
)

func TestInforderTraversal(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
		got := inorderTraversal(root)
		want := []int{1, 3, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test case 2", func(t *testing.T) {
		root := TreeNode{}.Right
		got := inorderTraversal(root)
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode{1, nil, nil}
		got := inorderTraversal(root)
		want := []int{1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
