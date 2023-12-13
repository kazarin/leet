package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextGreatestLetter(t *testing.T) {
	assert.Equal(t, byte('c'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	assert.Equal(t, byte('f'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
	assert.Equal(t, byte('x'), nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'))
	assert.Equal(t, byte('f'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
	assert.Equal(t, byte('c'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j'))
}
