/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
package main

// @lc code=start
func canChain(a, b byte) bool {
	if a == '1' {
		return true
	}
	if a == '2' && b <= '6' {
		return true
	}
	return false
}

func numDecodings(s string) int {
	var endNodes, chainNodes int

	switch s[0] {
	case '0':
		return 0
	case '1', '2':
		chainNodes = 1
		endNodes = 1
	default:
		chainNodes = 0
		endNodes = 1
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if chainNodes == 0 {
				return 0
			}
			chainNodes, endNodes = 0, chainNodes
		} else if s[i] == '1' || s[i] == '2' {
			chainNodes, endNodes = endNodes, endNodes+chainNodes
		} else if canChain(s[i-1], s[i]) {
			chainNodes, endNodes = 0, endNodes+chainNodes
		} else {
			chainNodes = 0
		}
	}
	return endNodes
}

// @lc code=end
