package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProductDifference(t *testing.T) {
	assert.Equal(t, 34, maxProductDifference([]int{5, 6, 2, 7, 4}))
	assert.Equal(t, 64, maxProductDifference([]int{4, 2, 5, 9, 7, 4, 8}))
}
