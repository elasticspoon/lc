package main

import "testing"

func TestTotalMoney(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{4, 10},
		{10, 37},
		{20, 96},
		{1000, 74926},
	}

	for _, tt := range tests {
		got := totalMoney(tt.n)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%d", tt, got)
		}
	}
}
