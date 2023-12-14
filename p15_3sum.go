package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				triplets = append(triplets, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}

	}
	return triplets
}
