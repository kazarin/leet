package main

func maxNumberOfBalloons(text string) int {
	storage := map[rune]int{}
	for _, token := range text {
		storage[token]++
	}

	balloons := 0
	for {
		for _, t := range "balloon" {
			if storage[t] == 0 {
				return balloons

			}
			storage[t]--
		}
		balloons++

	}

	//if storage['b'] == storage['a'] && storage['a'] == storage['l']/2 && storage['l']/2 == storage['o']/2 && storage['o']/2 == storage['n'] {
	//	return (storage['b'] + storage['a'] + storage['l']/2 + storage['o']/2 + storage['n']) / 5
	//}

	return balloons

}
