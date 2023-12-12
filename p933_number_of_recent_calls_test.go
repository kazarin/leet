package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_recentCounter(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, 1, obj.Ping(1))
	assert.Equal(t, 2, obj.Ping(3001))
	assert.Equal(t, 2, obj.Ping(3002))
}
