package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxScore(t *testing.T) {
	assert.Equal(t, 5, maxScore("011101"))
	assert.Equal(t, 5, maxScore("00111"))
	assert.Equal(t, 3, maxScore("1111"))
	assert.Equal(t, 1, maxScore("00"))
}
