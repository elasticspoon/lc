/*
 * @lc app=leetcode id=316 lang=typescript
 *
 * [316] Remove Duplicate Letters
 */

// @lc code=start
function removeDuplicateLetters(s: string): string {
  const charMap = new Map<string, number>();
  for (let i = 0; i < s.length; i++) {
    charMap.set(s[i], i);
  }

  const charStack: string[] = [];
  const charSet = new Set<string>();
  for (let i = 0; i < s.length; i++) {
    let char = s[i];
    if (charSet.has(char)) {
      continue;
    }

    while (charStack.length > 0) {
      let topChar = charStack.at(-1);

      if (topChar == undefined) return "error";
      if (char > topChar) break;

      let topCharIndex = charMap.get(topChar);
      if (topCharIndex === undefined) return "error";

      if (topCharIndex > i) {
        charStack.pop();
        charSet.delete(topChar);
      } else {
        break;
      }
    }
    charSet.add(char);
    charStack.push(char);
  }
  return charStack.join("");
}
// @lc code=end

console.log(removeDuplicateLetters("bca"));
console.log(removeDuplicateLetters("bcabc"));
console.log(removeDuplicateLetters("cbacdcbc"));
