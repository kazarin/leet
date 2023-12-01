package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	assert.Equal(t, true, hasCycle(IntsToListNodesWithCycle([]int{3, 0, 2, 4}, 1)))
	assert.Equal(t, true, hasCycle(IntsToListNodesWithCycle([]int{1, 2}, 0)))
	assert.Equal(t, false, hasCycle(IntsToListNodes([]int{1})))
}
