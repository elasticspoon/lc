/*
 * @lc app=leetcode id=1160 lang=golang
 *
 * [1160] Find Words That Can Be Formed by Characters
 */
package main

// @lc code=start
func countCharacters(words []string, chars string) int {
	charsCount := [26]int{}
	for i := 0; i < len(chars); i++ {
		charsCount[chars[i]%26]++
	}

	sum := 0
	for _, s := range words {
		countCopy := charsCount
		validWord := true
		for i := 0; i < len(s); i++ {
			countCopy[s[i]%26]--
			if countCopy[s[i]%26] < 0 {
				validWord = false
				break
			}
		}

		if validWord {
			sum += len(s)
		}
	}
	return sum
}

// @lc code=end
