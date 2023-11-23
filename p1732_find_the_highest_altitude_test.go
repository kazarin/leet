package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestAltitude(t *testing.T) {
	assert.Equal(t, 1, largestAltitude([]int{-5, 1, 5, 0, -7}))
	assert.Equal(t, 0, largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
