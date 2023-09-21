/*
 * @lc app=leetcode id=1658 lang=typescript
 *
 * [1658] Minimum Operations to Reduce X to Zero
 */

// @lc code=start
function minOperations(nums: number[], x: number): number {
  const reverseHash = new Map();
  reverseHash.set(0, 0);

  let accum = 0;
  for (let i = nums.length - 1; i >= 0; i--) {
    accum += nums[i];
    if (accum > x) {
      break;
    }

    reverseHash.set(accum, nums.length - i);
  }

  let minOps = Infinity;
  let forwardOps = 0;
  accum = 0;

  while (accum <= x && forwardOps < nums.length) {
    let reverseOps = reverseHash.get(x - accum);
    if (
      reverseOps != undefined &&
      reverseOps + forwardOps < minOps &&
      forwardOps + reverseOps <= nums.length
    ) {
      minOps = reverseOps + forwardOps;
    }
    accum += nums[forwardOps];
    forwardOps++;
  }

  return minOps == Infinity ? -1 : minOps;
}
// @lc code=end
