/*
 * @lc app=leetcode id=799 lang=typescript
 *
 * [799] Champagne Tower
 */

// @lc code=start
function champagneTower(
  poured: number,
  query_row: number,
  query_glass: number
): number {
  let avail = poured - (query_row * (query_row + 1)) / 2;

  console.log(avail);

  let res: number;

  if (query_row === 0) {
    res = avail;
  } else if (query_row === 1) {
    res = avail / 2;
  } else {
    res = avail / ((query_row - 1) * 2 + 2);
    if (query_glass !== 0 && query_glass !== query_row) {
      res = res * 2;
    }
  }

  return Math.min(Math.max(res, 0), 1);
}
// @lc code=end
