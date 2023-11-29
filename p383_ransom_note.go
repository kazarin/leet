package main

func canConstruct(ransomNote string, magazine string) bool {
	dictionary := map[rune]int{}

	for _, t := range magazine {
		dictionary[t]++
	}

	for _, n := range ransomNote {
		if dictionary[n] > 0 {
			dictionary[n]--
		} else {
			return false
		}
	}

	return true
}
