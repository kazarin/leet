package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	buffer := map[string][]string{}
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
		buffer[string(bytes)] = append(buffer[string(bytes)], str)

	}
	values := [][]string{}
	for _, vals := range buffer {
		values = append(values, vals)
	}
	return values

}
