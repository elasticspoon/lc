/*
 * @lc app=leetcode id=1359 lang=typescript
 *
 * [1359] Count All Valid Pickup and Delivery Options
 */

// @lc code=start
const MOD = 1_000_000_007;
function countOrders(n: number): number {
  let accum = 1;

  for (let i = 1; i <= n; i++) {
    accum = (accum * (i * 2 - 1) * i) % MOD;
  }

  return accum;
}
// @lc code=end
