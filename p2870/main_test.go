package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	assert.Equal(t, 4, minOperations([]int{2, 3, 3, 2, 2, 4, 2, 3, 4}))
	assert.Equal(t, -1, minOperations([]int{2, 1, 2, 2, 3, 3}))
	assert.Equal(t, 3, minOperations([]int{6, 6, 6, 6, 6, 5, 5}))
	assert.Equal(t, 7, minOperations([]int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}))
}
