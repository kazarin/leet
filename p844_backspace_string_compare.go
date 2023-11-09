package main

func backspaceCompare(s string, t string) bool {
	buf1 := []byte{}
	buf2 := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(buf1) >= 1 {
				buf1 = buf1[:len(buf1)-1]
			}
		} else {
			buf1 = append(buf1, s[i])
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			if len(buf2) >= 1 {
				buf2 = buf2[:len(buf2)-1]
			}
		} else {
			buf2 = append(buf2, t[i])
		}
	}

	return string(buf1) == string(buf2)
}
