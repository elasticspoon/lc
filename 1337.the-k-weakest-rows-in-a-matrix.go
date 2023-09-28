/*
 * @lc app=leetcode id=1337 lang=golang
 *
 * [1337] The K Weakest Rows in a Matrix
 */

package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func kWeakestRows(mat [][]int, k int) []int {
	strength := [][2]int{}
	for i, row := range mat {
		numSoldiers := 0
		for _, v := range row {
			if v == 0 {
				break
			}
			numSoldiers++
		}
		strength = append(strength, [2]int{numSoldiers, i})
	}

	sort.Slice(strength, func(i, j int) bool {
		if strength[i][0] == strength[j][0] {
			return strength[i][1] < strength[j][1]
		} else {
			return strength[i][0] < strength[j][0]
		}
	})

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, strength[i][1])
	}
	return res
}

// @lc code=end

// func main() {
// 	mat := [][]int{
// 		{1, 1, 0, 0, 0},
// 		{1, 1, 1, 1, 0},
// 		{1, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 0},
// 		{1, 1, 1, 1, 1},
// 	}
// 	fmt.Println(kWeakestRows(mat, 3))
// 	mat = [][]int{
// 		{1, 0, 0, 0},
// 		{1, 1, 1, 1},
// 		{1, 0, 0, 0},
// 		{1, 0, 0, 0},
// 	}
// 	fmt.Println(kWeakestRows(mat, 2))
// }
