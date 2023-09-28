/*
 * @lc app=leetcode id=316 lang=typescript
 *
 * [316] Remove Duplicate Letters
 */
// @lc code=start
function removeDuplicateLetters(s) {
    var charMap = new Map();
    for (var i = 0; i < s.length; i++) {
        charMap.set(s[i], i);
    }
    var charStack = [];
    var charSet = new Set();
    for (var i = 0; i < s.length; i++) {
        if (charSet.has(s[i])) {
        }
        else {
            charStack.push(s[i]);
            charSet.add(s[i]);
        }
    }
    return charStack.reverse().join("");
}
// @lc code=end
