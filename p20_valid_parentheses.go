package main

func isValid(s string) bool {
	buffer := []byte{}
	for i := 0; i < len(s); i++ {
		if len(buffer) > 0 && ((s[i] == ')' && buffer[len(buffer)-1] == '(') || (s[i] == '}' && buffer[len(buffer)-1] == '{') || (s[i] == ']' && buffer[len(buffer)-1] == '[')) {
			buffer = buffer[:len(buffer)-1]
		} else {
			buffer = append(buffer, s[i])
		}

	}

	if len(buffer) == 0 {
		return true
	} else {
		return false
	}
}
