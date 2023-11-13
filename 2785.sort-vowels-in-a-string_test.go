package main

import "testing"

func TestSortVowels(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"lEetcOde", "lEOtcede"},
		{"lYmpH", "lYmpH"},
		{"aA", "Aa"},
		{"a", "a"},
		{"kOOk", "kOOk"},
	}

	for _, tt := range tests {
		got := sortVowels(tt.in)
		if got != tt.out {
			t.Errorf("sortVowels(%v)=%v, want %v", tt.in, got, tt.out)
		}
	}
}
