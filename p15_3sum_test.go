package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.Equal(t, [][]int{}, threeSum([]int{0, 1, 1}))
	assert.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0}))
	assert.Equal(t, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
