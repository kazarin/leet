package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 2}))
	assert.Equal(t, 5, removeDuplicates([]int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	assert.Equal(t, 0, removeDuplicates([]int{}))
	assert.Equal(t, 1, removeDuplicates([]int{1, 1}))
}
