/*
 * @lc app=leetcode id=1048 lang=typescript
 *
 * [1048] Longest String Chain
 */

// @lc code=start
function longestStrChain(words: string[]): number {
  const wordArr = new Array<Map<string, number>>();
  for (let word of words) {
    const len = word.length;
    wordArr[len] = wordArr[len] || new Map<string, number>();
    wordArr[len].set(word, 1);
  }

  let max = 1;

  for (let i = 16; i > 0; i--) {
    if (!wordArr[i]) continue;
    for (let [word, pred] of wordArr[i]) {
      let potential_preds = word
        .split("")
        .map((_, i) => word.slice(0, i) + word.slice(i + 1));

      for (let potential_pred of potential_preds) {
        if (wordArr[i - 1]?.has(potential_pred)) {
          let x = wordArr[i - 1].get(potential_pred);
          x = Math.max(x, pred + 1);
          wordArr[i - 1].set(potential_pred, x);
          max = Math.max(max, x);
        }
      }
    }
  }

  return max;
}
// @lc code=end
