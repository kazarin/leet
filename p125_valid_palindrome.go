package main

func downcase(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c
	}
	return c + 32
}

func validCharacter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome2(s string) bool {
	left := 0
	buffer := []byte{}
	if len(s) == 1 {
		return true
	}
	for i := 0; i < len(s); i++ {
		if validCharacter(s[i]) {

			buffer = append(buffer, s[i])

		}
	}
	right := len(buffer) - 1
	for left < right {
		if downcase(buffer[left]) != downcase(buffer[right]) {
			return false
		}
		left++
		right--
	}
	return true

}
