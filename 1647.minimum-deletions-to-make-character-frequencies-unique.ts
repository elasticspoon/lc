/*
 * @lc app=leetcode id=1647 lang=typescript
 *
 * [1647] Minimum Deletions to Make Character Frequencies Unique
 */

// @lc code=start
function minDeletions(s: string): number {
  let histogram = {};
  let subtractions = 0;

  for (let i = 0; i < s.length; i++) {
    histogram[s.charCodeAt(i) % 26] =
      (histogram[s.charCodeAt(i) % 26] || 0) + 1;
  }

  let arr = [];
  for (let key in histogram) {
    arr[histogram[key]] = (arr[histogram[key]] || 0) + 1;
  }

  for (let i = arr.length - 1; i > 0; i--) {
    if (arr[i] > 1) {
      subtractions += arr[i] - 1;
      arr[i - 1] = (arr[i - 1] ?? 0) + (arr[i] - 1);
    }
  }

  return subtractions;
}
// @lc code=end
