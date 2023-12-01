package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	assert.ElementsMatch(t, []string{"AAAAACCCCC", "CCCCCAAAAA"}, findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	assert.ElementsMatch(t, []string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	assert.ElementsMatch(t, []string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAA"))
}
