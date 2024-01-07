package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	input1 := []int{0, 1, 0, 3, 12}
	moveZeroes(input1)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, input1)

	input2 := []int{0, 1}
	moveZeroes(input2)
	assert.Equal(t, []int{1, 0}, input2)
}
