package main

import "testing"

func TestWithinEditDist(t *testing.T) {
	tests := []struct {
		s1, s2 string
		n      int
		want   bool
	}{
		{"abc", "abc", 0, true},
		{"abc", "abc", 1, true},
		{"abcde", "abc", 1, false},
		{"abcde", "abc", 2, true},
		{"ab", "abc", 1, true},
		{"akc", "abc", 1, true},
	}

	for i, tt := range tests {
		got := withinEditDist(tt.s1, tt.s2, tt.n)
		if got != tt.want {
			t.Errorf("%d: withinEditDist(%s, %s, %d) = %t, want %t", i, tt.s1, tt.s2, tt.n, got, tt.want)
		}
	}
}
