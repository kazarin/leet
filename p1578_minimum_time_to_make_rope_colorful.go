package main

func minCost(colors string, neededTime []int) int {
	left := 0
	biggest := neededTime[0]
	sum := neededTime[0]
	out := 0
	for right := 1; right < len(colors); right++ {
		if colors[left] == colors[right] {
			sum += neededTime[right]
			biggest = max(biggest, neededTime[right])
		} else {
			out += sum - biggest
			left = right
			sum = neededTime[right]
			biggest = neededTime[right]
		}

	}

	out += sum - biggest

	return out
}
