package main

func numberOfBeams(bank []string) int {
	nums := []int{}
	for i := 0; i < len(bank); i++ {
		ones := 0
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				ones++

			}
		}
		nums = append(nums, ones)

	}
	out := 0
	for i := 0; i < len(nums); i++ {
		j := i - 1
		if i > 0 && nums[i] != 0 {
			for j > 0 && nums[j] == 0 {
				j--
			}
			out += nums[i] * nums[j]

		}
	}
	return out

}
