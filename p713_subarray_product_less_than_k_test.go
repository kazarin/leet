package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSubarrayProductLessThanK(t *testing.T) {
	assert.Equal(t, 8, numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	assert.Equal(t, 0, numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
	assert.Equal(t, 0, numSubarrayProductLessThanK([]int{1, 1, 1}, 1))
}
