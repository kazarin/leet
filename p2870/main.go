package main

func minOperations(nums []int) int {
	hash := map[int]int{}
	for _, n := range nums {
		hash[n]++
	}

	out := 0

	for _, v := range hash {
		if v < 2 {
			return -1
		}
		switch v % 3 {
		case 0:
			out += v / 3
		case 1:
			out += 2 + (v-4)/3
		case 2:
			out += 1 + (v-2)/3

		}
	}

	return out
}
