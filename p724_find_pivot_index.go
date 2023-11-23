package main

func pivotIndex(nums []int) int {
	sums := []int{}
	sums = append(sums, nums[0])
	for i := 1; i < len(nums); i++ {
		sums = append(sums, sums[i-1]+nums[i])

	}
	for i := 0; i < len(sums); i++ {
		left := 0
		right := 0
		if i == 0 {
			left = 0
		} else {
			left = sums[i-1]
		}
		if i == len(sums)-1 {
			right = 0
		} else {
			right = sums[len(sums)-1] - sums[i]
		}

		if left == right {
			return i
		}

	}
	return -1

}
