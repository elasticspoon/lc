/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package main

// @lc code=start

type palindrome struct {
	lettersOdd    map[byte]bool
	numOddLetters int
}

func (p *palindrome) isPalindrome() bool {
	return p.numOddLetters < 2
}

func (p *palindrome) toggleLetter(l byte) {
	if v, ok := p.lettersOdd[l]; ok {
		p.lettersOdd[l] = !v
		if v {
			p.numOddLetters--
		} else {
			p.numOddLetters++
		}
	}
}

func (p *palindrome) testLetter(l byte) bool {
	p.toggleLetter(l)
	defer p.toggleLetter(l)
	return p.isPalindrome()
}

//	func longestPalindrome(s string) string {
//		best := s[0:1]
//		left, right := 0, 1
//
//		pal := palindrome{}
//		pal.toggleLetter(s[0])
//		pal.toggleLetter(s[1])
//
//		for right < len(s) {
//			for pal.isPalindrome() {
//				if right < len(s)-1 && pal.testLetter(s[right+1]) {
//					right++
//					best = s[left : right+1]
//				} else if right < len(s)-1 && left > 0 && s[left-1] == s[right+1] {
//					right++
//					left--
//					best = s[left : right+1]
//				} else {
//					break
//				}
//			}
//			pal.toggleLetter(s[left])
//			left++
//			right++
//			pal.toggleLetter(s[right])
//
//		}
//
//		return best
//	}
func longestPalindrome(s string) string {
	best := s[0:1]
	left, right := 0, 0

	for right < len(s) {
		for {
			if right < len(s)-1 && isPalindrome(s[left:right+2]) {
				right++
				best = s[left : right+1]
			} else if right < len(s)-1 && left > 0 && isPalindrome(s[left-1:right+2]) {
				right++
				left--
				best = s[left : right+1]
			} else {
				break
			}
		}
		left++
		right++

	}

	return best
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2+1; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end
