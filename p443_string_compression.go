package main

import (
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	left, write := 0, 0
	for right := 1; right <= len(chars); right++ {
		if right == len(chars) || chars[right] != chars[left] {
			chars[write] = chars[left]
			write++
			count := right - left
			if count > 1 {
				ss := strconv.Itoa(count)
				for _, ch := range ss {
					chars[write] = byte(ch)
					write++
				}
			}
			left = right
		}
	}
	return write
}
