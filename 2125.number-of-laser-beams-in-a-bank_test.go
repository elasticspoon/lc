package main

import "testing"

func TestNumberOfBeams(t *testing.T) {
	tests := []struct {
		bank []string
		want int
	}{
		{[]string{"011001", "000000", "010100", "001000"}, 8},
		{[]string{"000", "111", "000"}, 0},
	}

	for _, tt := range tests {
		got := numberOfBeams(tt.bank)
		if got != tt.want {
			t.Errorf("numberOfBeams(%v) = %v, want %v", tt.bank, got, tt.want)
		}
	}
}
