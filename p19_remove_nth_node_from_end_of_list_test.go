package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	assert.Equal(t, IntsToListNodes([]int{2, 3, 4, 5, 6, 7, 8}), removeNthFromEnd(IntsToListNodes([]int{1, 2, 3, 4, 5, 6, 7, 8}), 8))
	assert.Equal(t, IntsToListNodes([]int{2, 3, 4, 5}), removeNthFromEnd(IntsToListNodes([]int{1, 2, 3, 4, 5}), 5))
	assert.Equal(t, IntsToListNodes([]int{1}), removeNthFromEnd(IntsToListNodes([]int{1, 2}), 1))
}
