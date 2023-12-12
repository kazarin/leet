package main

func findRow(matrix [][]int, target int) int {
	left := 0
	right := len(matrix) - 1

	for left <= right {
		mid := (left + right) / 2
		m := matrix[mid]
		if target >= m[0] && target <= m[len(m)-1] {
			return mid

		} else if target < m[0] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func searchMatrix(matrix [][]int, target int) bool {
	row := findRow(matrix, target)
	if row == -1 {
		return false
	}

	nums := matrix[row]

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		m := nums[mid]
		if target == m {
			return true

		} else if target < m {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
