package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCost(t *testing.T) {
	assert.Equal(t, 3, minCost("abaac", []int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, minCost("abc", []int{1, 2, 3}))
	assert.Equal(t, 2, minCost("aabaa", []int{1, 2, 3, 4, 1}))
	assert.Equal(t, 23, minCost("bbbaaa", []int{4, 9, 3, 8, 8, 9}))
	assert.Equal(t, 309, minCost("aababbaabababbbabbbaaaabaabbbabbbbaaab", []int{18, 5, 12, 7, 36, 33, 38, 18, 40, 2, 32, 14, 12, 15, 35, 33, 5, 28, 29, 6, 40, 28, 24, 28, 32, 27, 20, 11, 7, 34, 15, 21, 13, 26, 28, 39, 13, 16}))
}
