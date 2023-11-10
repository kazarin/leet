package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeGood(t *testing.T) {
	assert.Equal(t, "leetcode", makeGood("leEeetcode"))
	assert.Equal(t, "", makeGood("abBAcC"))
	assert.Equal(t, "s", makeGood("s"))
	assert.Equal(t, "", makeGood(""))
}
