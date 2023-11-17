package main

import (
	"strconv"
)

func calPoints(operations []string) int {
	buffer := []int{}
	for i := 0; i < len(operations); i++ {
		l := len(buffer) - 1

		switch op := operations[i]; op {
		case "+":
			buffer = append(buffer, buffer[l]+buffer[l-1])
		case "D":
			buffer = append(buffer, buffer[l]*2)
		case "C":
			buffer = buffer[:l]
		default:
			i, _ := strconv.Atoi(op)
			buffer = append(buffer, i)

		}
	}

	answer := 0
	for _, op := range buffer {
		answer += op
	}
	return answer
}
