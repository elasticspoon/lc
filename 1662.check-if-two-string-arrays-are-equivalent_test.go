package main

import "testing"

func TestArrayStringsAreEqual(t *testing.T) {
	tests := []struct {
		word1 []string
		word2 []string
		want  bool
	}{
		{[]string{"ab", "c"}, []string{"a", "bc"}, true},
		{[]string{"a", "cb"}, []string{"ab", "c"}, false},
		{[]string{"abc", "d", "defg"}, []string{"abcddefg"}, true},
	}

	for _, tt := range tests {
		got := arrayStringsAreEqual(tt.word1, tt.word2)
		if got != tt.want {
			t.Errorf("arrayStringsAreEqual(%v, %v) = %v, want %v", tt.word1, tt.word2, got, tt.want)
		}
	}
}
