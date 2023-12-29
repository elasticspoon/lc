package main

import "testing"

func TestGetLegthOfOptimalCompression(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"aaabcccd", 2, 4},
		{"aabbaa", 2, 2},
		{"aaaaaaaaaaa", 0, 3},
		{"llllllllllttttttttt", 1, 4},
	}

	for _, tt := range tests {
		got := getLengthOfOptimalCompression(tt.s, tt.k)
		if got != tt.want {
			t.Errorf("getLengthOfOptimalCompression(%q, %d) == %d, want %d", tt.s, tt.k, got, tt.want)
		}
	}
}
