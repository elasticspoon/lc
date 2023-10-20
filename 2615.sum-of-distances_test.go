package main

import "testing"

func TestDistance(t *testing.T) {
	tests := []struct {
		nums []int
		want []int64
	}{
		{[]int{1, 3, 1, 1, 2}, []int64{5, 0, 3, 4, 0}},
		{[]int{0, 5, 3}, []int64{0, 0, 0}},
		{[]int{2, 0, 2, 2, 6, 5, 2}, []int64{11, 0, 7, 7, 0, 0, 13}},
	}

	for _, tt := range tests {
		got := distance(tt.nums)
		if len(got) != len(tt.want) {
			t.Errorf("distance(%v) = %v; want %v", tt.nums, got, tt.want)
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("distance(%v) = %v; want %v", tt.nums, got, tt.want)
				break
			}
		}
	}
}
