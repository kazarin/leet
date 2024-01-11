package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST(t *testing.T) {
	assert.Equal(t, true, isValidBST(BinaryTree([]interface{}{2, 1, 3})))
	assert.Equal(t, false, isValidBST(BinaryTree([]interface{}{5, 1, 4, nil, nil, 3, 6})))
	assert.Equal(t, false, isValidBST(BinaryTree([]interface{}{2, 2, 2})))
}
