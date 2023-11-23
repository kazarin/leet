package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestUniqueNumber(t *testing.T) {
	assert.Equal(t, 8, largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1}))
	assert.Equal(t, -1, largestUniqueNumber([]int{9, 9, 8, 8}))
}
