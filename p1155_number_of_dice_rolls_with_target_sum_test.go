package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numRollsToTarget(t *testing.T) {
	assert.Equal(t, 4, numRollsToTarget(4, 4, 5))
	assert.Equal(t, 1, numRollsToTarget(1, 6, 3))
	assert.Equal(t, 2, numRollsToTarget(2, 5, 3))
	assert.Equal(t, 6, numRollsToTarget(2, 6, 7))
	assert.Equal(t, 222616187, numRollsToTarget(30, 30, 500))
}
