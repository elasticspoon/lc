/*
 * @lc app=leetcode id=880 lang=golang
 *
 * [880] Decoded String at Index
 */

package main

import (
	"fmt"
	"unicode"
)

// @lc code=start
type StringNode struct {
	size   int
	node   *StringNode
	string string
}

func (sn *StringNode) stringSize() int {
	return len(sn.string)
}

func decodeAtIndex(s string, k int) string {
	tree := &StringNode{0, nil, ""}

	index := 0
	for _, c := range s {
		if index >= k {
			break
		}
		if unicode.IsDigit(c) {
			size := int(c - '0')
			index *= size
			tree = &StringNode{index, tree, ""}
		} else {
			tree.string += string(c)
			tree.size++
			index++
		}
	}
	k -= 1

	for k >= 0 {
		if tree.stringSize() > 0 && k >= tree.size {
			k %= tree.size
		}
		if tree.stringSize() > 0 && k+tree.stringSize() >= tree.size {
			return string(tree.string[k-tree.size+tree.stringSize()])
		} else {
			tree = tree.node
		}
	}

	return ""
}

// @lc code=end
func main() {
	fmt.Println(decodeAtIndex("leet2code3", 10) == "o")
	fmt.Println(decodeAtIndex("ha22", 5) == "h")
	fmt.Println(decodeAtIndex("a2345678999999999999999", 1) == "a")
	fmt.Println(decodeAtIndex("a2b3c4d5e6f7g8h9", 10) == "c")
	fmt.Println(decodeAtIndex("vk6u5xhq9v", 554) == "k")
	fmt.Println(decodeAtIndex("y959q969u3hb22odq595", 222280369) == "y")
	fmt.Println(decodeAtIndex("ixm5xmgo78", 711) == "x")
}
