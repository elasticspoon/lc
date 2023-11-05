/*
 * @lc app=leetcode id=1441 lang=golang
 *
 * [1441] Build an Array With Stack Operations
 */
package main

// @lc code=start
func buildArray(target []int, n int) []string {
	result := []string{}

	streamVal := 1
	for _, v := range target {
		for {
			if streamVal > n {
				return result
			}
			result = append(result, "Push")
			if streamVal != v {
				result = append(result, "Pop")
			}
			streamVal++

			if result[len(result)-1] == "Push" {
				break
			}
		}
	}

	return result
}

// @lc code=end
