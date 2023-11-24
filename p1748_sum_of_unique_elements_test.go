package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumOfUnique(t *testing.T) {
	assert.Equal(t, 4, sumOfUnique([]int{1, 2, 3, 2}))
	assert.Equal(t, 0, sumOfUnique([]int{1, 1, 1, 1, 1}))
	assert.Equal(t, 15, sumOfUnique([]int{1, 2, 3, 4, 5}))
}
