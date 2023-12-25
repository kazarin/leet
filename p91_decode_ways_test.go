package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numDecodings(t *testing.T) {
	assert.Equal(t, 2, numDecodings("12"))
	assert.Equal(t, 3, numDecodings("226"))
	assert.Equal(t, 0, numDecodings("06"))
	assert.Equal(t, 1, numDecodings("27"))
}
