package main

import (
	"reflect"
	"testing"
)

func TestFindArray(t *testing.T) {
	tests := []struct {
		pref []int
		want []int
	}{
		{[]int{5, 2, 0, 3, 1}, []int{5, 7, 2, 3, 2}},
		{[]int{13}, []int{13}},
	}

	for _, tt := range tests {
		got := findArray(tt.pref)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findArray(%v) = %v, want %v", tt.pref, got, tt.want)
		}
	}
}
