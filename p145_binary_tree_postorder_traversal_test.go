package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_postorderTraversal(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, postorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
