package main

import (
	"reflect"
	"testing"
)

func TestInforderTraversal(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		root := &TreeNode94{1, nil, &TreeNode94{2, &TreeNode94{3, nil, nil}, nil}}
		got := inorderTraversal(root)
		want := []int{1, 3, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test case 2", func(t *testing.T) {
		root := TreeNode94{}.Right
		got := inorderTraversal(root)
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test case 3", func(t *testing.T) {
		root := &TreeNode94{1, nil, nil}
		got := inorderTraversal(root)
		want := []int{1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
