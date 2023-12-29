package main

func maxDepth(s string) int {
	stack := []byte{}
	out := 0

	for i := 0; i < len(s); i++ {
		l := len(stack)
		if s[i] == '(' {
			stack = append(stack, '(')
			l++
		} else if s[i] == ')' {
			out = max(l, out)
			stack = stack[:l-1]
			l--
		}

	}
	return out

}
