package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	assert.Equal(t, "blue is sky the", reverseWords("the sky is blue"))
	assert.Equal(t, "world hello", reverseWords("  hello world  "))
	assert.Equal(t, "example good a", reverseWords("a good   example"))

}
