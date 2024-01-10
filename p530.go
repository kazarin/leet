package main

import (
	"fmt"
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	var fx func(*TreeNode)
	var prevNode *TreeNode
	out := math.MaxInt
	fx = func(root *TreeNode) {
		if root == nil {
			return
		}
		fx(root.Left)
		if prevNode != nil {
			fmt.Printf("prev: %d, curr: %d\n", prevNode.Val, root.Val)
			out = min(out, abs(root.Val-prevNode.Val))
		}
		prevNode = root

		fx(root.Right)
	}
	fx(root)

	return out

}
