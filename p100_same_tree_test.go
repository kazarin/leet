package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSameTree(t *testing.T) {
	tree1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	tree4 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	assert.Equal(t, true, isSameTree(tree1, tree2))
	assert.Equal(t, false, isSameTree(tree3, tree4))
}
