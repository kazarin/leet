package main

import (
	"testing"
)

func Test_middleNode(t *testing.T) {
	tests := map[string]struct {
		input  *ListNode
		output *ListNode
	}{"example 1": {
		input:  IntsToListNodes([]int{1, 2, 3, 4, 5}),
		output: IntsToListNodes([]int{3, 4, 5})},
		"example 2": {input: IntsToListNodes([]int{1, 2, 3, 4, 5, 6}),
			output: IntsToListNodes([]int{4, 5, 6})}}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := middleNode(test.input), test.output; got.SameAs(expected) == false {
				t.Fatalf("got: %v, expected: %v", got, expected)
			}
		})
	}
}
