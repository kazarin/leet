package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	area := 0

	for left <= right {
		ls := min(height[left], height[right])
		area = max(ls*(right-left), area)
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}
	return area

}
