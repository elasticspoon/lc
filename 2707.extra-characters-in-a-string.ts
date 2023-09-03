/*
 * @lc app=leetcode id=2707 lang=typescript
 *
 * [2707] Extra Characters in a String
 *
 * https://leetcode.com/problems/extra-characters-in-a-string/description/
 *
 * algorithms
 * Medium (34.79%)
 * Likes:    1567
 * Dislikes: 69
 * Total Accepted:    56.1K
 * Total Submissions: 108.2K
 * Testcase Example:  '"leetscode"\n["leet","code","leetcode"]'
 *
 * You are given a 0-indexed string s and a dictionary of words dictionary. You
 * have to break s into one or more non-overlapping substrings such that each
 * substring is present in dictionary. There may be some extra characters in s
 * which are not present in any of the substrings.
 *
 * Return the minimum number of extra characters left over if you break up s
 * optimally.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetscode", dictionary = ["leet","code","leetcode"]
 * Output: 1
 * Explanation: We can break s in two substrings: "leet" from index 0 to 3 and
 * "code" from index 5 to 8. There is only 1 unused character (at index 4), so
 * we return 1.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: s = "sayhelloworld", dictionary = ["hello","world"]
 * Output: 3
 * Explanation: We can break s in two substrings: "hello" from index 3 to 7 and
 * "world" from index 8 to 12. The characters at indices 0, 1, 2 are not used
 * in any substring and thus are considered as extra characters. Hence, we
 * return 3.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 50
 * 1 <= dictionary.length <= 50
 * 1 <= dictionary[i].length <= 50
 * dictionary[i] and s consists of only lowercase English letters
 * dictionary contains distinct words
 *
 *
 */

function __minExtraChar(s: string, dictionary: string[]): number {
  let missed: number = 0;
  let max_reach: number = 0;
  let len: number = 0;

  for (let i = 0; i < s.length; i++) {
    for (let j = i + 1; j <= s.length; j++) {
      const substring = s.slice(i, j);
      // console.log(substring);
      if (dictionary.includes(substring)) {
        if (j >= max_reach) {
          if (substring.length > len) {
          }
          max_reach = j;
        }
        // console.log("found " + substring);
      }
    }
    if (max_reach <= i) {
      missed++;
      // console.log(max_reach + "missed\n");
    }
  }
  return missed;
}
class TrieNode {
  children: { [key: string]: TrieNode };
  isWord: boolean;

  constructor() {
    this.children = {};
    this.isWord = false;
  }
}

function buildTrie(dictionary: string[]): TrieNode {
  const root = new TrieNode();
  for (const word of dictionary) {
    let node = root;
    for (const ch of word) {
      if (!(ch in node.children)) {
        node.children[ch] = new TrieNode();
      }
      node = node.children[ch];
    }
    node.isWord = true;
  }
  return root;
}
function ______minExtraChar(s: string, dictionary: string[]): number {
  const root = buildTrie(dictionary);
  const n = s.length;
  const dp = new Array<number>(n + 1).fill(s.length + 1);
  dp[n] = 0; // No extra character for an empty string

  for (let start = n - 1; start >= 0; start--) {
    dp[start] = dp[start + 1] + 1; // Initialize with worst-case scenario
    let node = root;
    for (let end = start; end < n; end++) {
      if (!(s[end] in node.children)) {
        break;
      }
      node = node.children[s[end]];
      if (node.isWord) {
        dp[start] = Math.min(dp[start], dp[end + 1]);
      }
    }
  }

  return dp[0];
}
// @lc code=start

function minExtraChar(s: string, dictionary: string[]): number {
  const dp: number[] = Array(s.length + 1);
  const ds: Set<string> = new Set(dictionary);

  dp[0] = 0;

  for (let i = 1; i <= s.length; i++) {
    dp[i] = dp[i - 1] + 1;
    for (let j = 0; j <= i; j++) {
      if (ds.has(s.slice(j, i)) && dp[j] < dp[i]) {
        dp[i] = dp[j];
      }
    }
  }

  return dp[s.length];
}
// @lc code=end
