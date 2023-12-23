package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPathCrossing(t *testing.T) {
	assert.Equal(t, false, isPathCrossing("NES"))
	assert.Equal(t, true, isPathCrossing("NESWW"))
}
