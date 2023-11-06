package main

import "testing"

func TestLongestPalindromeSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"b", "b"},
		{"bb", "bb"},
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"jilkajsdflkjaksjdfkjsdkfjskdjfksjdfkjsdkfjsdkfjskjdfksjdfkjsracecarkasdjfkjadfkjasdkjfkasjdfkjasdfkjasdkfjaksdjfkasdjfkdjsfkjddfkjskdjfkajsdfkjadkfjaskdjfkasjdfkajsdfkjasdkfjasdkjfkasdjfkasjdfkjasdkfjasdkjfkasdjfkajsdfkjasdkjfaskdjfkasjdfkjasdkfjasdkfooiuqweroiuwqeoiurwoiueoqwiuroquweoriuqwoeiruweqoiurfoiqweuroiwqeuroiuqweoruqweoiruj", "racecar"},
	}

	for _, tt := range tests {
		actual := longestPalindrome(tt.input)
		if actual != tt.expected {
			t.Errorf("longestPalindrome(%s): expected '%s', actual '%s'", tt.input, tt.expected, actual)
		}
	}
}
