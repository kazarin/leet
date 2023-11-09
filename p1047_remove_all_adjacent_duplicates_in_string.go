package main

func removeDuplicates2(s string) string {
	stack := []rune{}
	str := []rune(s)
	i := 0
	for i < len(str) {
		if len(stack) > 0 && stack[len(stack)-1] == str[i] {
			stack = stack[:len(stack)-1]
			str = append(str[:i-1], str[i+1:]...)
			i = i - 1
		} else {
			stack = append(stack, str[i])
			i++
		}
	}

	return string(str)
}
