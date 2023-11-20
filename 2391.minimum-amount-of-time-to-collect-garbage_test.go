package main

import "testing"

func TestGarbageCollection(t *testing.T) {
	tests := []struct {
		garbage []string
		travel  []int
		want    int
	}{
		{[]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}, 21},
		{[]string{"MMM", "PGM", "GP"}, []int{3, 10}, 37},
		{[]string{"G", "G", "G", "G"}, []int{10, 10, 10}, 34},
		{[]string{"G", "G", "G", "GP"}, []int{10, 10, 10}, 65},
		{[]string{"M", "G", "G", "GP"}, []int{10, 10, 10}, 65},
		{[]string{"G", "M", "G", "GP"}, []int{10, 10, 10}, 75},
	}

	for _, tt := range tests {
		got := garbageCollection(tt.garbage, tt.travel)
		if got != tt.want {
			t.Errorf("garbageCollection(%v, %v) = %v, want %v", tt.garbage, tt.travel, got, tt.want)
		}
	}
}
