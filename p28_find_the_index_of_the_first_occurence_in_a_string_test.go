package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr(t *testing.T) {
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr("leetcode", "leeto"))
	assert.Equal(t, -1, strStr("", "str"))
	assert.Equal(t, 2, strStr("hello", "ll"))

}
