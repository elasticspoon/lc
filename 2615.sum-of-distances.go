/*
 * @lc app=leetcode id=2615 lang=golang
 *
 * [2615] Sum of Distances
 */
package main

// @lc code=start
func distance(nums []int) []int64 {
	numsMap := make(map[int][]int)
	result := make([]int64, len(nums))

	for i, v := range nums {
		numList, ok := numsMap[v]
		if ok {
			numsMap[v] = append(numList, i+numList[len(numList)-1])
		} else {
			numsMap[v] = []int{i}
		}
	}

	for _, indexs := range numsMap {
		l := len(indexs)

		for i := 0; i < l; i++ {
			current := indexs[i]
			if i > 0 {
				current -= indexs[i-1]
			}

			temp := 0
			if i > 0 {
				temp += (i * current) - indexs[i-1]
			}
			if i < l-1 {
				temp += indexs[l-1] - indexs[i] - ((l - i - 1) * current)
			}

			result[current] = int64(temp)

		}
	}

	return result
}

// @lc code=end
