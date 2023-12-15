package main

var memo map[int]int

func climbStairs(n int) int {
	a, b := 1, 1
	i := 0
	for i < n-1 {
		a, b = a+b, a
		i++
	}
	return a
}
