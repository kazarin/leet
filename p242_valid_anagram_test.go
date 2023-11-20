package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	assert.Equal(t, true, isAnagram("anagram", "nagaram"))
	assert.Equal(t, false, isAnagram("rat", "car"))
	assert.Equal(t, false, isAnagram("a", "ab"))
}
