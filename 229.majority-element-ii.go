/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 */

package main

import (
	"fmt"
)

// @lc code=start
// import "fmt"

func majorityElement(nums []int) []int {
	target := len(nums) / 3
	res := []int{}
	freqMap := make(map[int]int, 0)

	for _, num := range nums {
		freqMap[num]++
		if freqMap[num] == target+1 {
			res = append(res, num)
		}
	}
	return res
}

// @lc code=end
func main229() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{1, 2}))
}
