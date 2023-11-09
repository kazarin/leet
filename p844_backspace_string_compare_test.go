package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_backspaceCompare(t *testing.T) {
	assert.Equal(t, true, backspaceCompare("ab#c", "ad#c"))
	assert.Equal(t, true, backspaceCompare("ab##", "c#d#"))
	assert.Equal(t, true, backspaceCompare("a##c", "#a#c"))
}
