package main

import "testing"

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
	}
	for _, tc := range tests {
		if got := closeStrings(tc.word1, tc.word2); got != tc.want {
			t.Errorf("closeStrings(%v, %v) = %v, want %v", tc.word1, tc.word2, got, tc.want)
		}
	}
}
