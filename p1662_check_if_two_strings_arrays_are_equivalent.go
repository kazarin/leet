package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := ""
	s2 := ""
	for _, s := range word1 {
		s1 = s1 + s
	}
	for _, s := range word2 {
		s2 = s2 + s
	}
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
