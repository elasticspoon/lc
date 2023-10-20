/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 */
package main

// @lc code=start
func backspaceCompare(s string, t string) bool {
	sIndex := len(s) - 1
	tIndex := len(t) - 1

	for sSkips, tSkips := 0, 0; sIndex >= 0 || tIndex >= 0; sIndex, tIndex = sIndex-1, tIndex-1 {
		for tIndex >= 0 && (t[tIndex] == '#' || tSkips > 0) {
			if t[tIndex] == '#' {
				tSkips++
			} else {
				tSkips--
			}
			tIndex--
		}

		for sIndex >= 0 && (s[sIndex] == '#' || sSkips > 0) {
			if s[sIndex] == '#' {
				sSkips++
			} else {
				sSkips--
			}
			sIndex--
		}

		if sIndex == -1 || tIndex == -1 {
			return sIndex == tIndex
		} else if t[tIndex] != s[sIndex] {
			return false
		}

	}

	return sIndex == tIndex
}

// @lc code=end
