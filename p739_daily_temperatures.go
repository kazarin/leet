package main

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := []int{}
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			ls := len(stack) - 1
			r := stack[ls]
			answer[r] = i - r
			stack = stack[:ls]
		}
		stack = append(stack, i)
	}

	return answer
}
