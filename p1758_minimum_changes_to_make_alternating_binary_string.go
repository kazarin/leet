package main

func minOperations(s string) int {
	a := 0
	b := 0
	for i := 0; i < len(s); i++ {
		var n int
		if s[i] == '1' {
			n = 1
		} else {
			n = 0
		}
		if n == i%2 {
			a++
		} else {
			b++
		}
	}
	return min(a, b)
}
