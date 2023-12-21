/*
 * @lc app=leetcode id=2706 lang=golang
 *
 * [2706] Buy Two Chocolates
 */

package main

// @lc code=start
func buyChoco(prices []int, money int) int {
	var a, b int
	if prices[0] < prices[1] {
		a, b = prices[0], prices[1]
	} else {
		a, b = prices[1], prices[0]
	}

	for i := 2; i < len(prices); i++ {
		if prices[i] < a {
			b = a
			a = prices[i]
		} else if prices[i] < b {
			b = prices[i]
		}
	}

	if a+b <= money {
		return money - a - b
	} else {
		return money
	}
}

// @lc code=end
