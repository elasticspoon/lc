/*
 * @lc app=leetcode id=1326 lang=typescript
 *
 * [1326] Minimum Number of Taps to Open to Water a Garden
 *
 * https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/description/
 *
 * algorithms
 * Hard (47.34%)
 * Likes:    2985
 * Dislikes: 161
 * Total Accepted:    109.1K
 * Total Submissions: 212K
 * Testcase Example:  '5\n[3,4,1,1,0,0]'
 *
 * There is a one-dimensional garden on the x-axis. The garden starts at the
 * point 0 and ends at the point n. (i.e The length of the garden is n).
 *
 * There are n + 1 taps located at points [0, 1, ..., n] in the garden.
 *
 * Given an integer n and an integer array ranges of length n + 1 where
 * ranges[i] (0-indexed) means the i-th tap can water the area [i - ranges[i],
 * i + ranges[i]] if it was open.
 *
 * Return the minimum number of taps that should be open to water the whole
 * garden, If the garden cannot be watered return -1.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 5, ranges = [3,4,1,1,0,0]
 * Output: 1
 * Explanation: The tap at point 0 can cover the interval [-3,3]
 * The tap at point 1 can cover the interval [-3,5]
 * The tap at point 2 can cover the interval [1,3]
 * The tap at point 3 can cover the interval [2,4]
 * The tap at point 4 can cover the interval [4,4]
 * The tap at point 5 can cover the interval [5,5]
 * Opening Only the second tap will water the whole garden [0,5]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 3, ranges = [0,0,0,0]
 * Output: -1
 * Explanation: Even if you activate all the four taps you cannot water the
 * whole garden.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^4
 * ranges.length == n + 1
 * 0 <= ranges[i] <= 100
 *
 *
 */

// @lc code=start
function minTaps(n: number, ranges: number[]): number {
  const tuples: number[][] = [];
  for (let i = 0; i < ranges.length; i++) {
    let l = ranges[i];
    if (l == 0) {
      continue;
    }
    tuples.push([i - l, i + l]);
  }

  tuples.sort((a, b) => {
    return a[0] - b[0];
  });

  let taps: number = 0;
  let g_spot: number = 0;

  while (g_spot < n) {
    const valid_t = tuples.filter((tuple) => {
      return contains(g_spot, tuple) && tuple[1] > g_spot;
    });
    if (valid_t.length == 0) {
      return -1;
    }
    valid_t.sort((a, b) => {
      return b[1] - a[1];
    });
    taps += 1;
    g_spot = valid_t[0][1];
  }
  return taps;
}

function contains(v: number, tuple: number[]): boolean {
  return v >= tuple[0] && v <= tuple[1];
}
// @lc code=end
