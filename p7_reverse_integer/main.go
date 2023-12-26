package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true

	}
	n := strconv.Itoa(x)
	if negative {
		n = n[1:]
	}
	out := ""
	for _, s := range n {
		out = string(s) + out
	}
	result, _ := strconv.Atoi(out)
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	if negative {
		return -result
	} else {

		return result
	}
}
