/*
 * @lc app=leetcode id=823 lang=golang
 *
 * [823] Binary Trees With Factors
 */
package main

import (
	"fmt"
	"sort"
)

func stay() {
	fmt.Println("stay")
}

// @lc code=start
const modulo = 1000000007

func numFactoredBinaryTrees(arr []int) int {
	vals := make(map[int]int)
	sum := 0

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i, value := range arr {
		num := 1
		// fmt.Println("looping", vals, arr[i])

		for j := 0; j <= i; j++ {
			childA := arr[j]
			childB := value / childA

			// fmt.Println("checking", value, childA)
			if childA*childB != value {
				// fmt.Println("skipped")
				continue
			}

			if factor, ok := vals[childB]; ok {
				// fmt.Println("a: ", childA, "b: ", childB)
				num += factor * vals[childA]
				num %= modulo
			}

		}

		sum += num
		sum %= modulo
		vals[value] = num
	}

	// fmt.Println(vals)

	return sum
}

// @lc code=end
