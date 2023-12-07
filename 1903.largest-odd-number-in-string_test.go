package main

import "testing"

func TestLargestOddNumber(t *testing.T) {
	tests := []struct {
		num  string
		want string
	}{
		{"52", "5"},
		{"4206", ""},
		{"35427", "35427"},
		{"32782489638346578713315098393010310518347382", "327824896383465787133150983930103105183473"},
	}
	for _, tt := range tests {
		got := largestOddNumber(tt.num)
		if got != tt.want {
			t.Errorf("failed for %+v: got=%s", tt, got)
		}
	}
}
