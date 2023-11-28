/*
 * @lc app=leetcode id=2147 lang=golang
 *
 * [2147] Number of Ways to Divide a Long Corridor
 */
package main

// @lc code=start
const mod2147 = 1e9 + 7

func numberOfWays(corridor string) int {
	var consP, numS int
	numPoss := 1

	for _, c := range corridor {
		if c == 'S' {
			numS++
			if numS%2 == 0 {
				numPoss = ((consP + 1) * numPoss) % mod2147
				consP = 0
			}
		} else {
			if numS%2 == 0 && numS >= 2 {
				consP++
			}
		}
		// fmt.Printf("i: %d, c: %c, numPoss: %d,consP: %d, numS: %d\n", i, c, numPoss, consP, numS)
	}

	if numS%2 == 1 || numS == 0 {
		return 0
	}

	return numPoss
}

// @lc code=end
