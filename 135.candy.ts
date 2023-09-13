/*
 * @lc app=leetcode id=135 lang=typescript
 *
 * [135] Candy
 */

// @lc code=start
function candy(ratings: number[]): number {
  let sum = 1;
  let consInc = 0;
  let consDec = 0;
  let peak = 0;

  for (let i = 1; i < ratings.length; i++) {
    if (ratings[i] > ratings[i - 1]) {
      consDec = 0;
      consInc++;
      peak = consInc;
      sum += consInc + 1;
    } else if (ratings[i] == ratings[i - 1]) {
      consDec = 0;
      consInc = 0;
      peak = 0;
      sum++;
    } else {
      consInc = 0;
      consDec++;
      sum += consDec + (peak >= consDec ? 0 : 1);
    }
  }

  return sum;
}
// @lc code=end
