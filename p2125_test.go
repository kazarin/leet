package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfBeams(t *testing.T) {
	assert.Equal(t, 8, numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
	assert.Equal(t, 0, numberOfBeams([]string{"000", "111", "000"}))
	assert.Equal(t, 15, numberOfBeams([]string{"1111100", "0000000", "0000000", "1110000"}))
}
