package main

import (
	"strconv"
)

func toDigits(n int) (digits []int) {
	str := strconv.Itoa(n)
	digits = make([]int, len(str))

	for i, n := range str {
		digit, _ := strconv.Atoi(string(n))
		digits[i] = digit
	}
	return
}

func sumDigits(digits []int) (result int) {
	for _, digit := range digits {
		result += (digit * digit)
	}
	return
}

func isHappy(n int) bool {
	visited := map[int]int{}
	s := sumDigits(toDigits(n))
	for {
		if s == 1 {
			return true
		}
		if visited[s] != 0 {
			return false
		}
		visited[s]++

		s = sumDigits(toDigits(s))
	}
}
