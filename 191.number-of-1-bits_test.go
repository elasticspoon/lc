package main

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		num  uint32
		want int
	}{
		{0b00000000000000000000000000001011, 3},
		{0b00000000000000000000000010000000, 1},
	}

	for _, tt := range tests {
		got := hammingWeight(tt.num)
		if got != tt.want {
			t.Errorf("hammingWeight(%v) = %v, want %v", tt.num, got, tt.want)
		}
	}
}
