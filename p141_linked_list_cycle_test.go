package main

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	tests := map[string]struct {
		input  *ListNode
		output bool
	}{"example 1": {
		input:  IntsToListNodesWithCycle([]int{3, 0, 2, 4}, 1),
		output: true},
		"example 2": {
			input:  IntsToListNodesWithCycle([]int{1, 2}, 0),
			output: true,
		}, "example 3": {input: IntsToListNodes([]int{1}), output: false}, "example 4": {input: nil, output: false}}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := hasCycle(test.input), test.output; got != expected {
				t.Fatalf("got: %v, expected: %v", got, expected)
			}
		})
	}
}
