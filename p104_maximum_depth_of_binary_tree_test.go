package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	input1 := TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	input2 := TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	assert.Equal(t, 3, maxDepth(&input1))
	assert.Equal(t, 2, maxDepth(&input2))
}
