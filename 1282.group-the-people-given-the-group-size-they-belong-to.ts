/*
 * @lc app=leetcode id=1282 lang=typescript
 *
 * [1282] Group the People Given the Group Size They Belong To
 */

// @lc code=start
function groupThePeople(groupSizes: number[]): number[][] {
  let result: number[][] = [];
  let tmp = {};

  for (let i = 0; i < groupSizes.length; i++) {
    let arr = tmp[groupSizes[i]] || [];
    arr.push(i);
    if (arr.length == groupSizes[i]) {
      result.push(arr);
      tmp[groupSizes[i]] = [];
    } else {
      tmp[groupSizes[i]] = arr;
    }
  }

  return result;
}
// function groupThePeople(groupSizes: number[]): number[][] {
//   let result: number[][] = [];
//   let tmp: number[][] = [[]];
//
//   for (let i = 0; i < groupSizes.length; i++) {
//     let arr = tmp[groupSizes[i]] || [];
//     arr.push(i);
//     if (arr.length == groupSizes[i]) {
//       result.push(arr);
//       tmp[groupSizes[i]] = [];
//     } else {
//       tmp[groupSizes[i]] = arr;
//     }
//   }
//
//   return result;
// }
// @lc code=end
