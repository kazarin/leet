package main

func containsDuplicate(nums []int) bool {
	visited := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := visited[nums[i]]; ok {
			return true
		}
		visited[nums[i]] = i
	}

	return false
}
