package main

func plusOne(digits []int) []int {
	rem := 1

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		if digit+rem == 10 {
			digits[i] = 0
			rem = 1
		} else {
			digits[i] = digit + rem
			rem = 0
		}
	}
	if rem == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
