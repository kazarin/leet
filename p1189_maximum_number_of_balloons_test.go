package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxNumberOfBalloons(t *testing.T) {
	assert.Equal(t, 1, maxNumberOfBalloons("nlaebolko"))
	assert.Equal(t, 2, maxNumberOfBalloons("loonbalxballpoon"))
	assert.Equal(t, 0, maxNumberOfBalloons("leetcode"))
	assert.Equal(t, 1, maxNumberOfBalloons("balllllllllllloooooooooon"))
}
