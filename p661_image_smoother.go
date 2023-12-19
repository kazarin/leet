package main

func imageSmoother(img [][]int) [][]int {
	out := make([][]int, len(img))
	for i := 0; i < len(img); i++ {
		out[i] = make([]int, len(img[0]))
	}
	for i := range img {
		for j := range img[i] {
			values := []int{}
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 {
						continue
					}

					ni, nj := i+di, j+dj

					if ni >= 0 && ni < len(img) && nj >= 0 && nj < len(img[ni]) {
						values = append(values, img[ni][nj])
					}
				}
			}
			values = append(values, img[i][j])
			sum := 0
			for _, val := range values {
				sum += val
			}

			out[i][j] = sum / len(values)

		}
	}

	return out

}
