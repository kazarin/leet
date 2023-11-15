package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome2(t *testing.T) {
	assert.Equal(t, true, isPalindrome2("A man, a plan, a canal: Panama"))
	assert.Equal(t, false, isPalindrome2("race a car"))
	assert.Equal(t, true, isPalindrome2(" "))
	assert.Equal(t, false, isPalindrome2("0P"))
	assert.Equal(t, true, isPalindrome2("a"))
	assert.Equal(t, true, isPalindrome2("a."))
}
