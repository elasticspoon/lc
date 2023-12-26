package main

import (
	"fmt"
	"testing"
)

func TestNumRollsToTarget(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		target int
		want   int
	}{
		{1, 6, 3, 1},
		{2, 6, 7, 6},
		{2, 5, 10, 1},
		{1, 2, 3, 0},
		{30, 30, 500, 222616187},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d,%d", tt.n, tt.k, tt.target), func(t *testing.T) {
			if got := numRollsToTarget(tt.n, tt.k, tt.target); got != tt.want {
				t.Errorf("numRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
