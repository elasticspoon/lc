package main

import (
	"reflect"
	"testing"
)

func TestGetSumAbsoluteDifferences(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{2, 3, 5}, []int{4, 3, 5}},
		{[]int{1, 4, 6, 8, 10}, []int{24, 15, 13, 15, 21}},
	}
	for _, tt := range tests {
		got := getSumAbsoluteDifferences(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("getSumAbsoluteDifferences(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
