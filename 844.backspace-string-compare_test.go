package main

import "testing"

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"ab#c", "ad#c", true},
		{"a#c", "b", false},
		{"ab##", "c#d#", true},
		{"a##c", "#a#c", true},
		{"bxj##tw", "bxj###tw", false},
		{"nzp#o#g", "b#nzp#o#g", true},
	}

	for _, tt := range tests {
		if got := backspaceCompare(tt.s, tt.t); got != tt.want {
			t.Errorf("backspaceCompare(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
