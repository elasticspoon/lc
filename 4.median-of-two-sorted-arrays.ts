/*
 * @lc app=leetcode id=4 lang=typescript
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
  let len1 = nums1.length;
  let len2 = nums2.length;

  if (len1 > len2) {
    return findMedianSortedArrays(nums2, nums1);
  }

  let left = 0;
  let right = len1;
  let combinedMid = (len1 + len2 + 1) >> 1;

  while (left <= right) {
    let shortMid = left + ((right - left) >> 1);
    let longMid = combinedMid - shortMid;

    let maxLeft1 = shortMid === 0 ? -Infinity : nums1[shortMid - 1];
    let minRight1 = shortMid === len1 ? Infinity : nums1[shortMid];

    let maxLeft2 = longMid === 0 ? -Infinity : nums2[longMid - 1];
    let minRight2 = longMid === len2 ? Infinity : nums2[longMid];

    if (maxLeft1 <= minRight2 && maxLeft2 <= minRight1) {
      if ((len1 + len2) % 2 === 0) {
        return (
          (Math.max(maxLeft1, maxLeft2) + Math.min(minRight1, minRight2)) / 2
        );
      } else {
        return Math.max(maxLeft1, maxLeft2);
      }
    } else if (maxLeft1 > minRight2) {
      right = shortMid;
    } else {
      left = shortMid + 1;
    }
  }

  return 0;
}
// @lc code=end
