package main

import "testing"

func TestJobScheduling(t *testing.T) {
	cases := []struct {
		startTime []int
		endTime   []int
		profit    []int
		expect    int
	}{
		{
			startTime: []int{1, 2, 3, 3},
			endTime:   []int{3, 4, 5, 6},
			profit:    []int{50, 10, 40, 70},
			expect:    120,
		},
		{
			startTime: []int{1, 2, 3, 4, 6},
			endTime:   []int{3, 5, 10, 6, 9},
			profit:    []int{20, 20, 100, 70, 60},
			expect:    150,
		},
		{
			startTime: []int{1, 1, 1},
			endTime:   []int{2, 3, 4},
			profit:    []int{5, 6, 4},
			expect:    6,
		},
	}
	for _, c := range cases {
		got := jobScheduling(c.startTime, c.endTime, c.profit)
		if got != c.expect {
			t.Fatal(c, got)
		}
	}
}
