package main

import "testing"

func TestCountVowelPermutation(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{1, 5},
		{2, 10},
		{5, 68},
		{144, 18208803},
		{3, 19},
		{4, 35},
	}

	for _, test := range tests {
		ret := countVowelPermutation(test.input)
		if ret != test.output {
			t.Errorf("countVowelPermutation(%d); expected %d, got %d", test.input, test.output, ret)
		}
	}
}
