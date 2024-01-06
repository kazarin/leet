package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shuffle(t *testing.T) {
	assert.Equal(t, []int{2, 3, 5, 4, 1, 7}, shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	assert.Equal(t, []int{1, 4, 2, 3, 3, 2, 4, 1}, shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	assert.Equal(t, []int{1, 2, 1, 2}, shuffle([]int{1, 1, 2, 2}, 2))
}
