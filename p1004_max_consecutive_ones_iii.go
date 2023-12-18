package main

func longestOnes(nums []int, k int) int {
	current := 0
	answer := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			current++
		}

		for current > k {
			if nums[left] == 0 {
				current--
			}
			left++
		}
		answer = max(answer, right-left+1)

	}
	return answer
}
