/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func reverseWordsByterBuffer(s string) string {
	var out bytes.Buffer
	wordStack := []rune{}

	for _, letter := range s {
		if letter == rune(' ') {
			for len(wordStack) > 0 {
				out.WriteRune(wordStack[len(wordStack)-1])
				wordStack = wordStack[:len(wordStack)-1]
			}
			out.WriteString(" ")
		} else {
			wordStack = append(wordStack, letter)
		}
	}
	for len(wordStack) > 0 {
		out.WriteRune(wordStack[len(wordStack)-1])
		wordStack = wordStack[:len(wordStack)-1]
	}

	return out.String()
}

// @lc code=start
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		reversedWord := reverseString(word)
		words[i] = reversedWord
	}
	return strings.Join(words, " ")
}

func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

// @lc code=end
func main557() {
	input := "Let's take LeetCode contest"
	output := reverseWords(input)
	fmt.Println(output == "s'teL ekat edoCteeL tsetnoc")
	fmt.Println(output)
	input = "God Ding"
	output = reverseWords(input)
	fmt.Println(output == "doG gniD")
	fmt.Println(output)
}
