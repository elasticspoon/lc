package main

import (
	"reflect"
	"testing"
)

const null = -1 << 63

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode513
	}
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 3, 2, 5, 3, null, 9}, want: []int{1, 3, 9}},
		{input: []int{1, 2, 3}, want: []int{1, 3}},
		// {input: []int{1, null, 3, null, null, null, 4, null, null, 5}, want: []int{1, 3, 4, 5}},
	}

	for _, tt := range tests {
		tree := createTree(tt.input)
		result := largestValues(tree)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("largestValues(%v) = %v, want %v", tt.input, result, tt.want)
		}
	}
}

func createTree(list []int) *TreeNode513 {
	tree := &TreeNode513{Val: list[0]}
	queue := []*TreeNode513{tree}
	for i := 1; i < len(list); i++ {
		node := queue[0]
		queue = queue[1:]
		if list[i] != null {
			node.Left = &TreeNode513{Val: list[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(list) && list[i] != null {
			node.Right = &TreeNode513{Val: list[i]}
			queue = append(queue, node.Right)
		}
	}

	return tree
}
