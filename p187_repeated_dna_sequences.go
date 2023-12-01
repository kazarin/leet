package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	left := 0
	right := 10
	storage := map[string]int{}
	answer := []string{}
	for right <= len(s) {
		seq := s[left:right]
		storage[seq]++
		left++
		right++

	}
	for seq, c := range storage {
		if c > 1 {
			answer = append(answer, seq)
		}
	}

	return answer
}
