/*
 * @lc app=leetcode id=377 lang=typescript
 *
 * [377] Combination Sum IV
 */

// @lc code=start
function combinationSum4(nums: number[], target: number): number {
  let dp = [1];

  for (let i = 1; i <= target; i++) {
    dp[i] = 0;
    for (let j = 0; j < nums.length; j++) {
      if (i >= nums[j]) {
        dp[i] += dp[i - nums[j]];
      }
    }
  }

  return dp[target];
}
// @lc code=end
