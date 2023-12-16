package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"abc", "aabbcc", false},
		{"abc", "ab", false},
	}

	for _, tt := range tests {
		got := isAnagram(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("isAnagram(%s, %s): got %t want %t", tt.s, tt.t, got, tt.want)
		}
	}
}
