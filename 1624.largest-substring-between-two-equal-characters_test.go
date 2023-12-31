package main

import "testing"

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"aa", 0},
		{"abca", 2},
		{"cbzxy", -1},
		{"cabbac", 4},
	}
	for _, c := range cases {
		got := maxLengthBetweenEqualCharacters(c.input)
		if got != c.want {
			t.Errorf("maxLengthBetweenEqualCharacters(%q) == %d, want %d", c.input, got, c.want)
		}
	}
}
