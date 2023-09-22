/*
 * @lc app=leetcode id=392 lang=typescript
 *
 * [392] Is Subsequence
 */

// @lc code=start
function isSubsequence(s: string, t: string): boolean {
  let mainIndex = 0;

  for (let i = 0; i < s.length; i++) {
    let found = false;
    while (mainIndex < t.length) {
      if (s[i] == t[mainIndex]) {
        found = true;
        mainIndex++;
        break;
      } else {
        mainIndex++;
      }
    }
    if (!found) {
      return false;
    }
  }

  return true;
}
// @lc code=end
