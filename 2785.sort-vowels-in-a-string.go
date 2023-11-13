/*
 * @lc app=leetcode id=2785 lang=golang
 *
 * [2785] Sort Vowels in a String
 */
package main

import (
	"sort"
)

// @lc code=start
func sortVowels(s string) string {
	var vowels []rune
	var locs []int
	bytes := []byte(s)

	for i, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels = append(vowels, c)
			locs = append(locs, i)
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	for i := range vowels {
		bytes[locs[i]] = byte(vowels[i])
	}

	return string(bytes)
}

// @lc code=end
