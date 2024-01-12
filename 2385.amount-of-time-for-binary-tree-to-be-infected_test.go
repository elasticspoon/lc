package main

import "testing"

func TestAmountOfTime(t *testing.T) {
	tests := []struct {
		tree  []any
		start int
		want  int
	}{
		{[]any{1, 5, 3, nil, 4, 10, 6, 9, 2}, 3, 4},
		{[]any{1}, 1, 0},
	}

	for _, tt := range tests {
		tree := buildTree(tt.tree, 0)
		got := amountOfTime(tree, tt.start)
		if got != tt.want {
			t.Errorf("amountOfTime(%v, %v) = %v, want %v", tt.tree, tt.start, got, tt.want)
		}
	}
}
