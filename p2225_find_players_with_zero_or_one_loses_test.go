package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findWinners(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 10}, {4, 5, 7, 8}}, findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	assert.Equal(t, [][]int{{1, 2, 5, 6}, {}}, findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
}
