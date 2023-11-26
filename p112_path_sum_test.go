package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasPathSum(t *testing.T) {
	input1 := TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}}
	input2 := TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	assert.Equal(t, true, hasPathSum(&input1, 22))
	assert.Equal(t, false, hasPathSum(&input2, 5))
	assert.Equal(t, false, hasPathSum(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 1))
	assert.Equal(t, false, hasPathSum(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 0))
	assert.Equal(t, true, hasPathSum(&TreeNode{Val: -2, Right: &TreeNode{Val: -3}}, -5))
}
