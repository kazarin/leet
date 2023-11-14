package main

import "fmt"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	fmt.Printf("l: %d, r: %d\n", l, r)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
