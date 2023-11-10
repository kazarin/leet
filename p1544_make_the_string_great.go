package main

func makeGood(s string) string {
	buffer := []byte{}
	for i := 0; i < len(s); i++ {
		if len(buffer) > 0 {
			last := buffer[len(buffer)-1]
			if (s[i] >= 'a' && s[i] <= 'z' && s[i]-last == 32) || (s[i] >= 'A' && s[i] <= 'Z' && last-s[i] == 32) {
				buffer = buffer[:len(buffer)-1]
			} else {
				buffer = append(buffer, s[i])
			}
		} else {
			buffer = append(buffer, s[i])
		}
	}

	return string(buffer)
}
