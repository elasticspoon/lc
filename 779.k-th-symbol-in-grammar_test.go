package main

import "testing"

func TestKthGrammar(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{n: 1, k: 1, want: 0},
		{n: 2, k: 1, want: 0},
		{n: 2, k: 2, want: 1},
		{n: 2, k: 2, want: 1},
		{n: 5, k: 12, want: 1},
		{n: 5, k: 4, want: 0},
		{n: 4, k: 6, want: 0},
		{n: 4, k: 3, want: 1},
		{n: 30, k: 434991989, want: 0},
		{n: 30, k: 25, want: 0},
	}

	for _, tt := range tests {
		result := kthGrammar(tt.n, tt.k)
		if result != tt.want {
			t.Errorf("kthGrammar(%v, %v) = %v, want %v", tt.n, tt.k, result, tt.want)
		}
	}
}
