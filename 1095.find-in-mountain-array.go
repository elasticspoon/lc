/*
 * @lc app=leetcode id=1095 lang=golang
 *
 * [1095] Find in Mountain Array
 */
package main

// import "fmt"

type MountainArray struct {
	arr []int
}

func (m *MountainArray) get(index int) int {
	return m.arr[index]
}

func (m *MountainArray) length() int {
	return len(m.arr)
}

// @lc code=start
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	location := -1

	peak := findPeak(mountainArr)

	// searching on the left side of peak
	for l, r := 0, peak; l <= r; {
		c := (l + r) / 2
		curr := mountainArr.get(c)

		// fmt.Printf("l: %d, r: %d, c: %d, curr: %d target: %d\n", l, r, c, curr, target)
		if curr == target {
			return c
		} else if curr > target {
			r = c - 1
		} else {
			l = c + 1
		}
	}

	for l, r := peak, mountainArr.length()-1; l <= r; {
		c := (l + r) / 2
		curr := mountainArr.get(c)

		// fmt.Printf("l: %d, r: %d, c: %d, curr: %d target: %d\n", l, r, c, curr, target)
		if curr == target {
			return c
		} else if curr < target {
			r = c - 1
		} else {
			l = c + 1
		}
	}
	return location
}

func findPeak(mountainArr *MountainArray) int {
	n := mountainArr.length()

	for l, r := 0, n-1; l < r; {
		p := (l + r) / 2
		curr := mountainArr.get(p)
		left := mountainArr.get(p - 1)
		right := mountainArr.get(p + 1)

		// fmt.Printf("l: %d, r: %d, p: %d, curr: %d, left: %d, right: %d\n", l, r, p, curr, left, right)
		if curr > left && curr > right {
			return p
		} else if left > curr { // we are to the right of peak
			r = p
		} else if right > curr { // we are to the left of peak
			l = p
		}
	}

	return -1
}

// @lc code=end
