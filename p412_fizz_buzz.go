package main

import "strconv"

func fizzBuzz(n int) []string {
	answer := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			s := strconv.Itoa(i)
			answer = append(answer, s)
		}
	}
	return answer
}
