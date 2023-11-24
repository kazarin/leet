package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElements(t *testing.T) {
	input1 := IntsToListNodes([]int{1, 2, 6, 6, 3, 4, 5, 6, 6})
	output1 := IntsToListNodes([]int{1, 2, 3, 4, 5})

	assert.Equal(t, output1, removeElements(input1, 6))
	assert.Equal(t, &ListNode{}, removeElements(IntsToListNodes([]int{7, 7, 7, 7}), 7))
}
