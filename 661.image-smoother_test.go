package main

import (
	"reflect"
	"testing"
)

func TestImageSmoother(t *testing.T) {
	tests := []struct {
		img  [][]int
		want [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{[][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}, [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}}},
	}

	for _, tt := range tests {
		got := imageSmoother(tt.img)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("imageSmoother(%v) = %v, want %v", tt.img, got, tt.want)
		}
	}
}
