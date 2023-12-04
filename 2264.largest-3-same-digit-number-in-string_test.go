package main

import "testing"

func TestLargeGoodInteger(t *testing.T) {
	tests := []struct {
		num  string
		want string
	}{
		{"6777133339", "777"},
		{"42352338", ""},
		{"2300019", "000"},
	}

	for _, tt := range tests {
		got := largestGoodInteger(tt.num)
		if got != tt.want {
			t.Errorf("largestGoodInteger(%v) = %v, want %v", tt.num, got, tt.want)
		}
	}
}
