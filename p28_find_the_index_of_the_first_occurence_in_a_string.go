package main

func strStr(haystack string, needle string) int {
	left := 0
	right := len(needle)
	if len(haystack) == 0 {
		return -1
	}
	for {
		if right > len(haystack) {
			return -1
		}
		if haystack[left:right] == needle {
			return left
		} else {
			left++
			right++
		}

	}

	return 1
}
