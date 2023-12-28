package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compress(t *testing.T) {
	input1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	assert.Equal(t, 6, compress(input1))
	assert.Equal(t, []byte{'a', '2', 'b', '2', 'c', '3', 'c'}, input1)

	input2 := []byte{'a'}
	assert.Equal(t, 1, compress(input2))
	assert.Equal(t, []byte{'a'}, input2)

	input3 := []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}
	assert.Equal(t, 6, compress(input3))
	assert.Equal(t, []byte{'a', '3', 'b', '2', 'a', '2', 'a'}, input3)

}
