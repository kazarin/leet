package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMinimumDifference(t *testing.T) {
	assert.Equal(t, 1, getMinimumDifference(BinaryTree([]interface{}{4, 2, 6, 1, 3})))
	assert.Equal(t, 1, getMinimumDifference(BinaryTree([]interface{}{1, 0, 48, nil, nil, 12, 49})))
}
