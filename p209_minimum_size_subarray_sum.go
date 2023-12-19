package main

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	current := 0
	answer := math.MaxInt
	for right < len(nums) {
		current += nums[right]
		for current >= target {
			answer = min(answer, right-left+1)
			current -= nums[left]
			left++
		}
		right++
	}
	if right-left+1 > len(nums) {
		return 0
	}

	return answer
}
