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

	i, j := 0, len(nums)-1

	if nums == nil || j < 0 {
		return res
	}

	for i < j {
		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}

	if nums[i] != target {
		return res
	} else {
		res[0] = i
	}

	j = len(nums) - 1

	for i < j {
		mid := (i+j)/2 + 1
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid
		}
	}
	res[1] = j

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
