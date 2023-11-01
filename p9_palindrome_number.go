package main

import "fmt"

func isPalindrome(x int) bool {
	str := fmt.Sprint(x)
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] == str[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
