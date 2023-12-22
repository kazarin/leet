package main

import (
	"math"
)

func maxScore(s string) int {
	lsum := 0
	rsum := 0
	if s[0] == '0' {
		lsum++
	}

	for right := 1; right < len(s); right++ {
		if s[right] == '1' {
			rsum++
		}
	}

	out := math.MinInt

	for right := 1; right < len(s); right++ {
		if s[right] == '0' && right != len(s)-1 {
			lsum++
		}
		if s[right-1] == '1' && right != 1 {
			rsum--
		}
		out = max(out, rsum+lsum)
	}

	return out
}
