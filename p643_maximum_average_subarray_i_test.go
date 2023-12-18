package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxAverage(t *testing.T) {
	assert.Equal(t, 12.75, findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	assert.Equal(t, 5.0, findMaxAverage([]int{5}, 1))
}
