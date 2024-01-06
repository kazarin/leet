package main

func shuffle(nums []int, n int) []int {
	out := make([]int, len(nums))
	for i := 0; i < n; i++ {
		out[2*i] = nums[i]
		out[2*i+1] = nums[i+n]
	}
	return out

}
