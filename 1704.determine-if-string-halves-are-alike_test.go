package main

import "testing"

func TestHalvesAreAlike(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"book", true},
		{"textbook", false},
		{"MerryChristmas", false},
		{"AbCdEfGh", true},
		{"Uaec", false},
	}

	for _, tt := range tests {
		if got := halvesAreAlike(tt.s); got != tt.want {
			t.Errorf("halvesAreAlike(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
