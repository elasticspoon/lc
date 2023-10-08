/*
 * @lc app=leetcode id=1458 lang=golang
 *
 * [1458] Max Dot Product of Two Subsequences
 */

package main

import "fmt"

// @lc code=start
func maxDotProduct(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}

	dp[0][0] = nums1[0] * nums2[0]
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = max1458(dp[0][j-1], nums1[0]*nums2[j])
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = max1458(dp[i-1][0], nums1[i]*nums2[0])
		for j := 1; j < len(dp[0]); j++ {

			newV := nums1[i] * nums2[j]
			newV = max1458(dp[i-1][j-1]+newV, newV)
			newV = max1458(newV, dp[i][j-1])
			dp[i][j] = max1458(dp[i-1][j], newV)
		}

	}

	return dp[len(nums1)-1][len(nums2)-1]
}

func max1458(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
func main1458() {
	nums1 := []int{2, 1, -2, 5}
	nums2 := []int{3, 0, -6}
	fmt.Println(maxDotProduct(nums1, nums2) == 18)

	nums1 = []int{3, -2}
	nums2 = []int{2, -6, 7}
	fmt.Println(maxDotProduct(nums1, nums2) == 21)

	nums1 = []int{3, -2, 6, 8}
	nums2 = []int{-7, 4, -8}
	fmt.Println(maxDotProduct(nums1, nums2) == 46)

	nums1 = []int{-1, -1}
	nums2 = []int{1, 1}
	fmt.Println(maxDotProduct(nums1, nums2) == -1)

	nums1 = []int{-3, -8, 3, -10, 1, 3, 9}
	nums2 = []int{9, 2, 3, 7, -9, 1, -8, 5, -1, -1}
	fmt.Println(maxDotProduct(nums1, nums2) == 200)

	nums1 = []int{-3, -8}
	nums2 = []int{9, 2, 3, 7, -9}
	fmt.Println(maxDotProduct(nums1, nums2) == 72)
}

func printDP1458(dp [][]int) {
	for _, row := range dp {
		fmt.Println(row)
	}
}
