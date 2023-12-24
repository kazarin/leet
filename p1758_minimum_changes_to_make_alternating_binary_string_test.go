package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	assert.Equal(t, 1, minOperations("0100"))
	assert.Equal(t, 0, minOperations("01"))
	assert.Equal(t, 2, minOperations("1111"))
}
