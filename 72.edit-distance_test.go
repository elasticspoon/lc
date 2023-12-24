package main

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{
		// {"horse", "ros", 3},
		// {"intention", "execution", 5},
		// {"", "", 0},
		{"zoologicoarchaeologist", "zoogeologist", 10},
		{"z", "zoogeologist", 11},
		{"zo", "zoogeologist", 10},
	}

	for _, tt := range tests {
		got := minDistance(tt.word1, tt.word2)
		if got != tt.want {
			t.Errorf("minDistance(%v, %v)=%v, want %v", tt.word1, tt.word2, got, tt.want)
		}
	}
}
