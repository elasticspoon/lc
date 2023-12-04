/*
 * @lc app=leetcode id=2264 lang=golang
 *
 * [2264] Largest 3-Same-Digit Number in String
 */
package main

// @lc code=start
func largestGoodInteger(num string) string {
	subs := ""

	for i := 0; i < len(num)-2; {
		if num[i] != num[i+1] {
			i++
			continue
		}

		if num[i] != num[i+2] {
			i += 2
			continue
		}

		if subs < num[i:i+3] {
			subs = num[i : i+3]
			if subs == "999" {
				return subs
			}

		}
		i += 3
	}

	return subs
}

// @lc code=end
