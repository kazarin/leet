package main

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1

	for left < right {
		mid := (left + right) / 2
		if target < letters[mid] {
			right = mid
		} else {
			left = mid + 1

		}
	}
	if letters[left] <= target {
		return letters[0]
	} else {
		return letters[left]
	}

}
