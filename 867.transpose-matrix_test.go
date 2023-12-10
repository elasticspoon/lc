package main

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	}

	for _, tt := range tests {
		got := transpose(tt.matrix)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
