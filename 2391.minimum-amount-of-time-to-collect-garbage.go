/*
 * @lc app=leetcode id=2391 lang=golang
 *
 * [2391] Minimum Amount of Time to Collect Garbage
 */

package main

// @lc code=start
func garbageCollection(garbage []string, travel []int) int {
	time := 0
	gMax, mMax, pMax := 0, 0, 0

	for i, g := range garbage {
		time += len(g)
		for _, c := range g {
			switch c {
			case 'G':
				gMax = i
			case 'M':
				mMax = i
			case 'P':
				pMax = i
			}
		}
	}

	for i, acc := 0, 0; i < len(travel); i++ {
		acc += travel[i]
		if i == gMax-1 {
			time += acc
		}
		if i == mMax-1 {
			time += acc
		}
		if i == pMax-1 {
			time += acc
		}

	}

	return time
}

// @lc code=end
