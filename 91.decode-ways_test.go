package main

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"0", 0},
		{"06", 0},
		{"10", 1},
		{"1040", 0},
		{"27", 1},
		{"2611055971756562", 4},
		{"2611055971", 2},
		{"26110559717", 4},
		{"2222", 5},
	}

	for _, tt := range tests {
		got := numDecodings(tt.s)
		if got != tt.want {
			t.Errorf("numDecodings(%v)=%v, want %v", tt.s, got, tt.want)
		}
	}
}
