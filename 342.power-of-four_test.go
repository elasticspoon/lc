package main

import "testing"

func TestPowerofFour(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{16, true},
		{3, false},
		{4, true},
		{5, false},
		{0, false},
		{4194304, true},
	}

	for _, tt := range tests {
		got := isPowerOfFour(tt.n)
		if got != tt.want {
			t.Errorf("isPowerOfFour(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
