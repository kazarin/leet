package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)

	nums2 := []int{1}
	merge(nums2, 1, []int{}, 0)

	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
	assert.Equal(t, []int{1}, nums2)
}
