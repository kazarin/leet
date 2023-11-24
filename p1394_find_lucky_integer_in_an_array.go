package main

func findLucky(arr []int) int {
	hash := map[int]int{}
	for _, n := range arr {
		hash[n]++
	}

	l := -1
	for k, v := range hash {
		if k == v {
			l = max(l, k)
		}
	}
	return l

}
