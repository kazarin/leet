package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for i < len(nums)-1 {
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
		i++
	}
	return len(nums)

}
