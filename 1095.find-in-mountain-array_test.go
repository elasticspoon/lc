/*
 * @lc app=leetcode id=1095 lang=golang
 *
 * [1095] Find in Mountain Array
 */
package main

import (
	"testing"
)

func TestFindInMR(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 3, 1}, 3, 2},
		{[]int{0, 1, 2, 4, 2, 1}, 3, -1},
		{[]int{1, 5, 2}, 2, 2},
		{[]int{1, 5, 2}, 5, 1},
		{[]int{1, 5, 2}, 1, 0},
		{[]int{1, 5, 2}, 0, -1},
		{[]int{1, 5, 2}, 6, -1},
		{[]int{1, 5, 2}, 4, -1},
		{[]int{3, 5, 3, 2, 0}, 0, 4},
		{[]int{3, 4, 7, 8, 9, 10, 11, 12, 11}, 8, 3},
		{[]int{3, 5, 3, 2, 0}, 3, 0},
	}

	for _, tt := range tests {
		ma := &MountainArray{tt.arr}
		if actual := findInMountainArray(tt.target, ma); actual != tt.expected {
			t.Errorf("findInMountainArray(%v, %v): expected %v, actual %v\n", tt.target, tt.arr, tt.expected, actual)
		}
	}
}

func TestFindPeak(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 3, 1}, 4},
		{[]int{3, 4, 7, 8, 9, 10, 11, 12, 11}, 7},
		{[]int{3, 5, 3, 2, 0}, 1},
	}

	for _, tt := range tests {
		ma := &MountainArray{tt.arr}
		actual := findPeak(ma)
		if actual != tt.expected {
			t.Errorf("findPeak(%v): expected %v, actual %v\n", tt.arr, tt.expected, actual)
		}
	}
}
