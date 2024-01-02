/*
 * @lc app=leetcode id=2610 lang=golang
 *
 * [2610] Convert an Array Into a 2D Array With Conditions
 */

package main

// @lc code=start
// func findMatrix(nums []int) [][]int {
// 	sort.Ints(nums)
//
// 	result := [][]int{{nums[0]}}
//
// 	for i, consValues := 1, 0; i < len(nums); i++ {
//
// 		if nums[i] == nums[i-1] {
// 			consValues++
// 		} else {
// 			consValues = 0
// 		}
//
// 		if consValues >= len(result) {
// 			result = append(result, []int{})
// 		}
// 		result[consValues] = append(result[consValues], nums[i])
// 	}
//
// 	return result
// }

func findMatrix(nums []int) [][]int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	result := [][]int{}
	// for i := 0; i < maxDepth; i++ {
	// 	result = append(result, []int{})
	// }

	for k, v := range count {
		for i := 0; i < v; i++ {
			if i >= len(result) {
				result = append(result, []int{})
			}
			result[i] = append(result[i], k)
		}
	}

	return result
}

// @lc code=end
