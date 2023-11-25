package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		switch token {
		case "+":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], a+b)
		case "*":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], a*b)
		case "/":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], b/a)
		case "-":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], b-a)

		default:
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)

		}
	}
	return stack[0]
}
