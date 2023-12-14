package main

import (
	"reflect"
	"testing"
)

func TestOnesMinusZeros(t *testing.T) {
	tests := []struct {
		grid [][]int
		want [][]int
	}{
		{[][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}, [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}}},
		{[][]int{{1, 1, 1}, {1, 1, 1}}, [][]int{{5, 5, 5}, {5, 5, 5}}},
	}

	for _, tt := range tests {
		got := onesMinusZeros(tt.grid)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
