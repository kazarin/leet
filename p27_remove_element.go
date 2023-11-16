package main

func removeElement(nums []int, val int) int {
	answer := 0
	right := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for nums[right] == val && right > i {
				right--
			}
			nums[i], nums[right] = nums[right], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			answer++
		}
	}
	return answer

}
