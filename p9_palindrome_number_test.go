package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input  int
	output bool
}{
	{121, true},
	{-121, false},
	{10, false},
	{1, true},
	{313, true}}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			s := isPalindrome(tt.input)
			if s != tt.output {
				t.Errorf("got: %v, want: %v", s, tt.output)
			}
		})

	}
}
