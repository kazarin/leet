package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preorderTraversal(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, preorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
