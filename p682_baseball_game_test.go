package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calPoints(t *testing.T) {
	assert.Equal(t, 30, calPoints([]string{"5", "2", "C", "D", "+"}))
	assert.Equal(t, 27, calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))

}
