package main

func isPathCrossing(path string) bool {
	start := [2]int{0, 0}
	visited := map[[2]int]bool{}
	visited[[2]int{0, 0}] = true
	for _, s := range path {
		switch s {
		case 'N':
			start[1]++
		case 'W':
			start[0]--
		case 'E':
			start[0]++
		case 'S':
			start[1]--
		}
		if visited[start] {
			return true
		}
		visited[start] = true
	}

	return false
}
