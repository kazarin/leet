package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLucky(t *testing.T) {
	assert.Equal(t, 2, findLucky([]int{2, 2, 3, 4}))
	assert.Equal(t, 3, findLucky([]int{1, 2, 2, 3, 3, 3}))
	assert.Equal(t, -1, findLucky([]int{2, 2, 2, 3, 3}))

}
