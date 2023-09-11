/*
 * @lc app=leetcode id=1282 lang=golang
 *
 * [1282] Group the People Given the Group Size They Belong To
 */

package main

// @lc code=start
func groupThePeople(groupSizes []int) [][]int {
	sizeMap := make(map[int][]int)

	res := [][]int{}
	for i, v := range groupSizes {
		sizeMap[v] = append(sizeMap[v], i)
		if len(sizeMap[v]) == v {
			res = append(res, sizeMap[v])
			sizeMap[v] = []int{}
		}
	}

	return res
}

// @lc code=end
