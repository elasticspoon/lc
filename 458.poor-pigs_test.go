package main

import "testing"

func TestPoorPigs(t *testing.T) {
	tests := []struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
		expected      int
	}{
		{1000, 15, 60, 5},
		{4, 15, 15, 2},
		{4, 15, 30, 2},
		{4, 15, 45, 1},
		{16, 15, 15, 4},
		{16, 15, 30, 3},
		{16, 15, 45, 2},
		{16, 1, 15, 1},
		{16, 1, 14, 2},
	}

	for _, test := range tests {
		ret := poorPigs(test.buckets, test.minutesToDie, test.minutesToTest)
		if ret != test.expected {
			t.Errorf("poorPigs(%d, %d, %d); expected %d, got %d", test.buckets, test.minutesToDie, test.minutesToTest, test.expected, ret)
		}
	}
}
