package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {

	assert.Equal(t, IntsToListNodes([]int{5, 4, 3, 2, 1}), reverseList(IntsToListNodes([]int{1, 2, 3, 4, 5})))
	//assert.Equal(t, IntsToListNodes([]int{2, 1}), reverseList(IntsToListNodes([]int{1, 2})))
}
