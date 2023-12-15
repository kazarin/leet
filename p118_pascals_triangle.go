package main

func generate(numRows int) [][]int {
	out := make([][]int, 2)
	out[0] = make([]int, 1)
	out[1] = make([]int, 2)
	if numRows == 1 {
		return [][]int{{1}}
	}

	out[0][0] = 1
	out[1][0] = 1
	out[1][1] = 1
	for i := 2; i < numRows; i++ {
		values := make([]int, i+1)
		values[0] = 1
		values[i] = 1
		for j := 1; j < i; j++ {
			values[j] = out[i-1][j-1] + out[i-1][j]
		}
		out = append(out, values)
	}
	return out
}
