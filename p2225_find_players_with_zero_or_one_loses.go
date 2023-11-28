package main

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {
	w := []int{}
	l := []int{}
	winners := map[int]int{}
	losers := map[int]int{}
	for _, match := range matches {
		winner := match[0]
		loser := match[1]
		winners[winner]++
		losers[loser]++
	}

	for winner, _ := range winners {
		if losers[winner] == 0 {
			w = append(w, winner)
		}
	}
	for loser, v := range losers {
		if v == 1 {
			l = append(l, loser)
		}
	}
	answer := make([][]int, 2)
	sort.Ints(w)
	sort.Ints(l)
	answer[0] = w
	answer[1] = l

	return answer

}
