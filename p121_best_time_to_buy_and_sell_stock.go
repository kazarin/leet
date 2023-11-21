package main

func maxProfit(prices []int) int {
	left := 0
	right := 1
	current := 0

	for right < len(prices) {
		for prices[left] > prices[right] {
			left++
		}
		current = max(current, prices[right]-prices[left])
		right++
	}

	return current
}
