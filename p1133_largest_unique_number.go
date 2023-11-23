package main

func largestUniqueNumber(nums []int) int {
	buffer := map[int]int{}
	for _, n := range nums {
		buffer[n]++
	}
	largest := -1
	for n, times := range buffer {
		if times == 1 {
			largest = max(largest, n)
		}

	}
	return largest

}
