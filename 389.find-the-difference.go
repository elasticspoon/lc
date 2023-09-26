/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */
package main

// @lc code=start
//
//	func findTheDifference(s string, t string) byte {
//		hist := [26]int{}
//		for i := 0; i < len(s); i++ {
//			hist[s[i]%26]++
//		}
//
//		for i := 0; i < len(t); i++ {
//			index := t[i] % 26
//			if hist[index] == 0 {
//				return t[i]
//			}
//			hist[index]--
//		}
//		return t[0]
//	}
func findTheDifference(s string, t string) byte {
	var cumDiff byte
	for i := 0; i < len(s); i++ {
		cumDiff += t[i] - s[i]
	}

	return cumDiff + t[len(s)]
}

// @lc code=end

// func main() {
// 	fmt.Println(findTheDifference("abcd", "abcde") == 'e')
// 	fmt.Println(findTheDifference("", "y") == 'y')
// 	fmt.Println(findTheDifference("a", "aa") == 'a')
// 	fmt.Println(findTheDifference("ae", "aea") == 'a')
// }
