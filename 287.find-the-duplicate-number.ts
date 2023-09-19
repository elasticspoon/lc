/*
 * @lc app=leetcode id=287 lang=typescript
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
// function findDuplicate(nums: number[]): number {
//   for (let i = 0; i < nums.length; i++) {
//     for (let j = 0; j < nums.length; j++) {
//       if (i === j) {
//         continue;
//       }
//       if (nums[i] === nums[j]) {
//         return nums[i];
//       }
//     }
//   }
// }
function findDuplicate(nums: number[]): number {
  let listCopy = new Array(nums.length);
  for (let n of nums)
    if (listCopy[n] == undefined) {
      listCopy[n] = true;
    } else {
      return n;
    }
  return -1;
}
// function findDuplicate(nums: number[]): number {
//   let listCopy = new Array(10_001);
//   for (let n of nums)
//     if (listCopy[n] == undefined) {
//       listCopy[n] = true;
//     } else {
//       return n;
//     }
//   return -1;
// }
// @lc code=end
