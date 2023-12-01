package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isHappy(t *testing.T) {
	assert.Equal(t, true, isHappy(19))
	assert.Equal(t, false, isHappy(2))
	assert.Equal(t, true, isHappy(7))
}
