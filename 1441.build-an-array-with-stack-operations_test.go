package main

import (
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	tests := []struct {
		target []int
		n      int
		want   []string
	}{
		{[]int{1, 3}, 3, []string{"Push", "Push", "Pop", "Push"}},
		{[]int{1, 2, 3}, 3, []string{"Push", "Push", "Push"}},
		{[]int{1, 2}, 4, []string{"Push", "Push"}},
		{[]int{2, 3, 4}, 4, []string{"Push", "Pop", "Push", "Push", "Push"}},
	}

	for _, tt := range tests {
		got := buildArray(tt.target, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("buildArray(%v, %d) = %v, want %v", tt.target, tt.n, got, tt.want)
		}
	}
}
