package main

func lengthOfLIS(nums []int) int {
	out := 1
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		j := i + 1
		m := dp[i]
		for ; j < len(nums); j++ {
			if nums[j] > nums[i] {
				m = max(m, dp[j])
			}

		}
		dp[i] = m + 1
		out = max(dp[i], out)
	}

	return out
}
