package main

func findMaxAverage(nums []int, k int) float64 {
	current := 0
	for i := 0; i < k; i++ {
		current += nums[i]
	}
	var answer float64 = float64(current) / float64(k)
	for i := k; i < len(nums); i++ {
		current += nums[i] - nums[i-k]
		answer = max(answer, float64(current)/float64(k))
	}

	return answer
}
