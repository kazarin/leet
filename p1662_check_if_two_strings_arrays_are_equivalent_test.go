package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_arrayStringsAreEqual(t *testing.T) {
	assert.Equal(t, true, arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	assert.Equal(t, false, arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	assert.Equal(t, true, arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}
