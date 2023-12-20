package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buyChoco(t *testing.T) {
	assert.Equal(t, 0, buyChoco([]int{1, 2, 2}, 3))
	assert.Equal(t, 3, buyChoco([]int{3, 2, 3}, 3))
}
