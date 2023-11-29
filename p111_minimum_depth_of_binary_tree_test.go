package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDepth(t *testing.T) {
	assert.Equal(t, 2, minDepth(&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}))
	assert.Equal(t, 5, minDepth(&TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}))
}
