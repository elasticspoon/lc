/*
 * @lc app=leetcode id=338 lang=typescript
 *
 * [338] Counting Bits
 *
 * https://leetcode.com/problems/counting-bits/description/
 *
 * algorithms
 * Easy (76.32%)
 * Likes:    10027
 * Dislikes: 454
 * Total Accepted:    889.8K
 * Total Submissions: 1.2M
 * Testcase Example:  '2'
 *
 * Given an integer n, return an array ans of length n + 1 such that for each i
 * (0 <= i <= n), ans[i] is the number of 1's in the binary representation of
 * i.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: [0,1,1]
 * Explanation:
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 *
 *
 * Example 2:
 *
 *
 * Input: n = 5
 * Output: [0,1,1,2,1,2]
 * Explanation:
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 * 3 --> 11
 * 4 --> 100
 * 5 --> 101
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 10^5
 *
 *
 *
 * Follow up:
 *
 *
 * It is very easy to come up with a solution with a runtime of O(n log n). Can
 * you do it in linear time O(n) and possibly in a single pass?
 * Can you do it without using any built-in function (i.e., like
 * __builtin_popcount in C++)?
 *
 *
 */

// @lc code=start
function countBits(n: number): number[] {
  const bits: number[] = [0];

  for (let i = 0; i <= n; i += 2) {
    // bits[i] = bits[Math.floor(i / 2)] + (i % 2);
    // idk if a call to math floor is avoidable here. JK we found it
    const v = bits[i >> 1];
    bits[i] = v;
    bits[i + 1] = v + 1;
  }

  return bits.slice(0, n + 1);
}
// @lc code=end
