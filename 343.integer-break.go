/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

package main

// @lc code=start
import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	m3s := n / 3
	r3s := n % 3

	inter := math.Pow(float64(3), float64(m3s))

	if r3s == 2 {
		inter *= 2
	} else if r3s == 1 {
		inter = inter / 3 * 4
	}

	return int(inter)
}

// @lc code=end
func main343() {
	fmt.Println(integerBreak(2) == 1)
	fmt.Println(integerBreak(3) == 2)
	fmt.Println(integerBreak(4) == 4)
	fmt.Println(integerBreak(5) == 6)
	fmt.Println(integerBreak(6) == 9)
	fmt.Println(integerBreak(7) == 12)
	fmt.Println(integerBreak(8) == 18)
	fmt.Println(integerBreak(9) == 27)
	fmt.Println(integerBreak(10) == 36)
	fmt.Println(integerBreak(11) == 54)
	fmt.Println(integerBreak(12) == 81)
	fmt.Println(integerBreak(58) == 1549681956)
}
