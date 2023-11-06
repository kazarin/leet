package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	assert.Equal(t, isValid("()"), true)
	assert.Equal(t, isValid("()[]{}"), true)
	assert.Equal(t, isValid("(]"), false)
	assert.Equal(t, isValid(""), true)
}
