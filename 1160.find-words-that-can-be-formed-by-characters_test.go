package main

import "testing"

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		words []string
		chars string
		want  int
	}{
		{[]string{"cat", "bt", "hat", "tree"}, "atach", 6},
		{[]string{"hello", "world", "leetcode"}, "welldonehoneyr", 10},
	}

	for _, tt := range tests {
		got := countCharacters(tt.words, tt.chars)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d\n", tt, got)
		}
	}
}
