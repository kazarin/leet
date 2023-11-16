package main

func getConcatenation(nums []int) []int {
	l := len(nums) * 2
	result := make([]int, l)
	for i := 0; i < len(nums); i++ {
		result[i] = nums[i]
		result[i+len(nums)] = nums[i]
	}
	return result

}
