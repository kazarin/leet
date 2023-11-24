package main

func sumOfUnique(nums []int) int {
	hash := map[int]int{}
	for _, n := range nums {
		hash[n]++
	}

	sum := 0
	for _, n := range nums {
		if hash[n] == 1 {
			sum += n
		}
	}

	return sum
}
