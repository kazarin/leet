package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 3}, twoSum([]int{2, 3, 4}, 6))
	assert.Equal(t, []int{1, 2}, twoSum([]int{-1, 0}, -1))
}
