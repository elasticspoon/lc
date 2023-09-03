/*
 * @lc app=leetcode id=2707 lang=golang
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
package main

// @lc code=start
func minExtraChar(s string, dictionary []string) int {
	dp := make([]int, len(s)+1)
	trie := trieData()

	for _, wr := range dictionary {
		trie.insert(wr)
	}
	// ds := make(map[string]bool, len(dictionary))
	// for _, v := range dictionary {
	// 	ds[v] = true
	// }

	for i := 1; i <= len(s); i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j <= i; j++ {
			// if conts(dictionary, s[j:i]) && dp[j] < dp[i] {
			if trie.search(s[j:i]) && dp[j] < dp[i] {
				// if ds[s[j:i]] && dp[j] < dp[i] {
				dp[i] = dp[j]
			}
		}
	}

	return dp[len(s)]
}

// Declaring trie_Node  for creating node in a trie
type trie_Node struct {
	// assigning limit of 26 for child nodes
	childrens [26]*trie_Node
	// declaring a bool variable to check the word end.
	wordEnds bool
}

// Initializing the root of the trie
type trie struct {
	root *trie_Node
}

// inititlaizing a new trie
func trieData() *trie {
	t := new(trie)
	t.root = new(trie_Node)
	return t
}

// Passing words to trie
func (t *trie) insert(word string) {
	current := t.root
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = new(trie_Node)
		}
		current = current.childrens[index]
	}
	current.wordEnds = true
}

// Initializing the search for word in node
func (t *trie) search(word string) bool {
	current := t.root
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	if current.wordEnds {
		return true
	}
	return false
}

// initializing the main function

// func conts(dict []string, s string) bool {
// 	for i := 0; i < len(dict); i++ {
// 		if dict[i] == s {
// 			return true
// 		}
// 	}
// 	return false
// }

// @lc code=end
