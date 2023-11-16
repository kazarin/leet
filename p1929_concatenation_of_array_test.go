package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getConcatenation(t *testing.T) {
	assert.Equal(t, []int{1, 2, 1, 1, 2, 1}, getConcatenation([]int{1, 2, 1}))
	assert.Equal(t, []int{1, 3, 2, 1, 1, 3, 2, 1}, getConcatenation([]int{1, 3, 2, 1}))
}
