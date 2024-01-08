package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rangeSumBST(t *testing.T) {
	assert.Equal(t, 22, rangeSumBST(&TreeNode{Val: 10, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 30}}, 5, 15))
}
