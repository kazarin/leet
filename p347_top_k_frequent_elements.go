package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}
	for _, n := range nums {
		hash[n]++
	}

	keys := []int{}
	for k, _ := range hash {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool { return hash[keys[i]] > hash[keys[j]] })

	return keys[:k]
}
