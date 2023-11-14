/*
 * @lc app=leetcode id=1930 lang=golang
 *
 * [1930] Unique Length-3 Palindromic Subsequences
 */
package main

// @lc code=start
func countPalindromicSubsequence(s string) int {
	var count int
	letters := make([][]int, 26)

	for i := 0; i < len(s); i++ {
		letters[s[i]%26] = append(letters[s[i]%26], i)
	}

	for _, indexs := range letters {
		if len(indexs) < 2 {
			continue
		}

		startI, endI := indexs[0], indexs[len(indexs)-1]

		if endI-startI < 2 {
			continue
		}

		for _, midIs := range letters {
			if len(midIs) == 0 {
				continue
			}

			for _, v := range midIs {
				if v > startI && v < endI {
					count++
					break
				}
			}
		}
	}

	return count
}

// @lc code=end
