package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	answer := 0
	product := 1
	left := 0
	for right := 0; right < len(nums); right++ {
		product = product * nums[right]
		for product >= k {
			product = product / nums[left]
			left++
		}

		answer += right - left + 1
	}

	return answer

}
