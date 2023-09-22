/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

package main

import (
	"math"
)

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	left, right := 0, len1
	combinedMid := (len1 + len2 + 1) / 2

	for left <= right {
		shortMid := left + (right-left)/2
		longMid := combinedMid - shortMid

		var minRight1, minRight2, maxLeft1, maxLeft2 int
		if shortMid == 0 {
			maxLeft1 = -math.MaxInt32
		} else {
			maxLeft1 = nums1[shortMid-1]
		}
		if shortMid == len1 {
			minRight1 = math.MaxInt32
		} else {
			minRight1 = nums1[shortMid]
		}

		if longMid == 0 {
			maxLeft2 = -math.MaxInt32
		} else {
			maxLeft2 = nums2[longMid-1]
		}
		if longMid == len2 {
			minRight2 = math.MaxInt32
		} else {
			minRight2 = nums2[longMid]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (len2+len1)%2 == 0 {
				return (math.Max(float64(maxLeft1), float64(maxLeft2)) + math.Min(float64(minRight1), float64(minRight2))) / 2.0
			} else {
				return math.Max(float64(maxLeft1), float64(maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			right = shortMid
		} else {
			left = shortMid + 1
		}
	}
	return 0
}

// @lc code=end

// func main() {
// 	arr1 := []int{1, 3}
// 	arr2 := []int{2}
//
// 	fmt.Println(findMedianSortedArrays(arr1, arr2))
// }
