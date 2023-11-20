package main

func isAnagram(s string, t string) bool {
	buffer := map[string]int{}
	for i := 0; i < len(s); i++ {
		t := string(s[i])
		buffer[t]++
	}

	for i := 0; i < len(t); i++ {
		t := string(t[i])
		if _, ok := buffer[t]; ok {
			buffer[t]--
			if buffer[t] == 0 {
				delete(buffer, t)
			}
		} else {
			buffer[t] = 1
		}
	}

	if len(buffer) == 0 {
		return true
	} else {
		return false
	}
}
