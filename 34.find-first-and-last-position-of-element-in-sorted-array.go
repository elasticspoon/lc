/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
package main

import "fmt"

// @lc code=start
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	l := len(nums)
	if l == 0 {
		return res
	}

	for left, right, p := 0, l-1, l/2; left <= right; {
		if (p == 0 && nums[p] == target) ||
			(p > 0 && nums[p-1] != target && nums[p] == target) {
			res[0] = p
			break
		}

		if nums[p] == target || nums[p] > target {
			right = p - 1
			p = (left + p) / 2
		} else {
			left = p + 1
			p = (p + right + 1) / 2
		}
	}

	for left, right, p := 0, l-1, l/2; left <= right; {
		if (p == (l-1) && nums[p] == target) ||
			(p < l-1 && nums[p+1] != target && nums[p] == target) {
			res[1] = p
			break
		}

		if nums[p] == target || nums[p] < target {
			left = p + 1
			p = (p + right + 1) / 2
		} else {
			right = p - 1
			p = (left + p) / 2
		}
	}

	return res
}

// @lc code=end

func main34() {
	fmt.Println(arraysIsEqual(searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4}))
	fmt.Println(arraysIsEqual(searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1}))
	fmt.Println(arraysIsEqual(searchRange([]int{8, 8, 8, 8, 8, 8, 8, 8}, 8), []int{0, 7}))
	fmt.Println(arraysIsEqual(searchRange([]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 3), []int{2, 5}))
	fmt.Println(arraysIsEqual(searchRange([]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 1), []int{0, 0}))
	fmt.Println(arraysIsEqual(searchRange([]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 9), []int{8, 8}))
	fmt.Println(arraysIsEqual(searchRange([]int{}, 0), []int{-1, -1}))
	fmt.Println(arraysIsEqual(searchRange([]int{1}, 0), []int{-1, -1}))
}

func arraysIsEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
