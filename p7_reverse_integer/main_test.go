package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
	assert.Equal(t, 21, reverse(120))
	assert.Equal(t, 0, reverse(1534236469))

}
