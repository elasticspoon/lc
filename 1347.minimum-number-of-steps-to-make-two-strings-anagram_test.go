package main

import "testing"

func TestMinSteps(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want int
	}{
		{"bab", "aba", 1},
		{"leetcode", "practice", 5},
		{"anagram", "mangaar", 0},
	}

	for _, tt := range tests {
		got := minSteps(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("minSteps(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
