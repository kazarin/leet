package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates2(t *testing.T) {
	assert.Equal(t, "ca", removeDuplicates2("abbaca"))
	assert.Equal(t, "ay", removeDuplicates2("azxxzy"))
}
