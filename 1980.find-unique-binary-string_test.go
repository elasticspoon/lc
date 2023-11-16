package main

import "testing"

func TestFindDiffBinaryString(t *testing.T) {
	tests := []struct {
		nums []string
	}{
		{[]string{"01", "10"}},
		{[]string{"00", "01"}},
		{[]string{"111", "011", "001"}},
		{[]string{"0011", "0111", "0110", "1011"}},
	}

	for _, tt := range tests {
		got := findDifferentBinaryString(tt.nums)
		if len(got) != len(tt.nums[0]) {
			t.Errorf("findDifferentBinaryString(%v) = %v, want %v", tt.nums, got, tt.nums[0])
		}
		for _, n := range tt.nums {
			if got == n {
				t.Errorf("findDifferentBinaryString(%v) = %v, want %v", tt.nums, got, n)
			}
		}

	}
}
