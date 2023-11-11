/*
 * @lc app=leetcode id=1743 lang=golang
 *
 * [1743] Restore the Array From Adjacent Pairs
 */

package main

// @lc code=start
func restoreArray(adjacentPairs [][]int) []int {
	result := []int{}

	adjacentMap := make(map[int][]int)
	for _, pair := range adjacentPairs {
		adjacentMap[pair[0]] = append(adjacentMap[pair[0]], pair[1])
		adjacentMap[pair[1]] = append(adjacentMap[pair[1]], pair[0])
	}

	current := 0
	for k, v := range adjacentMap {
		if len(v)%2 == 1 {
			current = k
			break
		}
	}

	for len(result) < len(adjacentPairs)+1 {
		result = append(result, current) // add current to result;

		for i := 0; i < len(adjacentMap[current]); i++ {
			if len(result) > 1 && adjacentMap[current][i] == result[len(result)-2] && len(adjacentMap[current]) > 1 {
				adjacentMap[current][i], adjacentMap[current][len(adjacentMap[current])-1] = adjacentMap[current][len(adjacentMap[current])-1], adjacentMap[current][i]
				adjacentMap[current] = adjacentMap[current][:len(adjacentMap[current])-1]
			}
		}

		if len(adjacentMap[current]) == 0 {
			break
		}
		current, adjacentMap[current] = adjacentMap[current][0], adjacentMap[current][1:]

	}

	return result
}

// @lc code=end
