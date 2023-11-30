package main

import (
	"strings"
)

func reverse(left, right int, runes []rune) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

}

func reverseWords(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i, _ := range s {
		if !(s[i] == 32 && i+1 < len(s) && s[i+1] == 32) {
			b.WriteRune(rune(s[i]))
		}
	}
	s = b.String()

	s = strings.TrimSpace(s)
	runes := []rune(s)
	left := 0
	right := len(runes) - 1
	reverse(left, right, runes)

	left = 0
	right = 0
	for left != len(runes) && right != len(runes) {
		if runes[right] == 32 {
			reverse(left, right-1, runes)
			left = right + 1
		}
		if right == len(s)-1 {
			reverse(left, right, runes)
			left = right + 1
		}
		right++
	}
	return string(runes)
}
