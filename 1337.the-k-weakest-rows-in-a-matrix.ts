/*
 * @lc app=leetcode id=1337 lang=typescript
 *
 * [1337] The K Weakest Rows in a Matrix
 */

// @lc code=start
function kWeakestRows(mat: number[][], k: number): number[] {
  const strengths: number[][] = [[]];
  for (let i = 0; i < mat.length; i++) {
    let sum = mat[i].reduce((acc, item) => {
      return (acc += item);
    });
    strengths[i] = [sum, i];
  }

  strengths.sort((a, b) => {
    if (a[0] == b[0]) {
      return a[1] - b[1];
    } else {
      return a[0] - b[0];
    }
  });

  return strengths.slice(0, k).map((item) => {
    return item[1];
  });
}
// @lc code=end
