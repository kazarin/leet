package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	assert.Equal(t, 3, maxDepth("(1+(2*3)+((8)/4))+1"))
	assert.Equal(t, 3, maxDepth("(1)+((2))+(((3)))"))
	assert.Equal(t, 1, maxDepth("()"))
}
