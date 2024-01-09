package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_leafSimilar(t *testing.T) {
	assert.Equal(t, true, leafSimilar(BinaryTree([]interface{}{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}), BinaryTree([]interface{}{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8})))
	assert.Equal(t, false, leafSimilar(BinaryTree([]interface{}{1, 2, 3}), BinaryTree([]interface{}{1, 3, 2})))
}
