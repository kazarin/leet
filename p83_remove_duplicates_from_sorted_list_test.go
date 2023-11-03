package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	input1 := IntsToListNodes([]int{1, 1, 2})
	input2 := IntsToListNodes([]int{1, 1, 2, 3, 3})

	assert.Equal(t, deleteDuplicates(input1).ToArray(), []int{1, 2})
	assert.Equal(t, deleteDuplicates(input2).ToArray(), []int{1, 2, 3})
}
