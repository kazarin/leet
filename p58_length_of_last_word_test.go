package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 4, lengthOfLastWord("   fly me   to   the moon  "))
	assert.Equal(t, 6, lengthOfLastWord("luffy is still joyboy"))
	assert.Equal(t, 1, lengthOfLastWord(" a"))
	assert.Equal(t, 1, lengthOfLastWord("a "))
	assert.Equal(t, 3, lengthOfLastWord("day"))
}
