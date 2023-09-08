/*
 * @lc app=leetcode id=118 lang=typescript
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
function generate(numRows: number): number[][] {
  const triangle = [[1]];

  for (let i = 1; i < numRows; i++) {
    const prevRow = triangle[i - 1];
    const row = [1];
    let j = 1;

    for (; j < prevRow.length; j++) {
      row[j] = prevRow[j - 1] + prevRow[j];
    }
    row[j] = 1;
    triangle[i] = row;
  }

  return triangle;
}
// @lc code=end
