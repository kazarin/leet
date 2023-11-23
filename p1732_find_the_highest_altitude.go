package main

func largestAltitude(gain []int) int {
	sums := []int{}
	sums = append(sums, gain[0])
	for i := 1; i < len(gain); i++ {
		sums = append(sums, sums[i-1]+gain[i])

	}
	highest := 0
	for _, n := range sums {
		highest = max(highest, n)
	}
	return highest

}
