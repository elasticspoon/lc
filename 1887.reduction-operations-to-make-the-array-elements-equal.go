/*
 * @lc app=leetcode id=1887 lang=golang
 *
 * [1887] Reduction Operations to Make the Array Elements Equal
 */

package main

// @lc code=start
func reductionOperations(nums []int) int {
	// freq := make([]int, 10)
	freq := make([]int, 5e4+1)

	for _, v := range nums {
		freq[v]++
	}

	sum := 0
	smallest := 0

	for i := 0; i < len(freq); i++ {
		if freq[i] > 0 {
			smallest = i
			break
		}
	}

	// fmt.Println(freq)
	for tot, i := 0, len(freq)-1; i > smallest; i-- {
		if freq[i] > 0 {
			tot += freq[i]
			sum += tot
		}
	}
	// fmt.Println(freq)

	return sum
}

// @lc code=end
