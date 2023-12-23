package main

import "math"

func withinEditDist(s1, s2 string, n int) bool {
	if math.Abs(float64(len(s1)-len(s2))) > float64(n) {
		return false
	}

	left, right, dist := 0, 0, 0

	for left < len(s1) && right < len(s2) {
		if s1[left] == s2[right] {
			left++
			right++
			continue
		}

		dist++
		if dist > n {
			return false
		}

		if left+1 >= len(s1) || right+1 >= len(s2) {
			return true
		}

		if s1[left+1] == s2[right] {
			left++
		} else if s1[left] == s2[right+1] {
			right++
		} else {
			left++
			right++
		}

	}
	return dist <= n
}
