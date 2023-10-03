/*
 * @lc app=leetcode id=1512 lang=golang
 *
 * [1512] Number of Good Pairs
 */
package main

// import "fmt"

// @lc code=start
func numIdenticalPairs(nums []int) int {
	pairMap := make(map[int]int, 1)

	for i := 0; i < len(nums); i++ {
		pairMap[nums[i]]++
	}

	result := 0
	for _, v := range pairMap {
		result += ((v - 1) * v) / 2
	}

	return result
}

// @lc code=end
// func main1512() {
// 	nums := []int{1, 2, 3, 1, 1, 3}
// 	fmt.Println(numIdenticalPairs(nums) == 4)
// 	nums = []int{1, 1, 1, 1}
// 	fmt.Println(numIdenticalPairs(nums) == 6)
// 	nums = []int{1, 2, 3}
// 	fmt.Println(numIdenticalPairs(nums) == 0)
// }
