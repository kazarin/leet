package main

import "sort"

func isValidBST(root *TreeNode) bool {
	values := []int{}
	var fx func(*TreeNode)
	fx = func(root *TreeNode) {
		if root == nil {
			return
		}
		fx(root.Left)
		values = append(values, root.Val)
		fx(root.Right)
	}
	fx(root)
	sorted := make([]int, len(values))
	copy(sorted, values)
	sort.Ints(sorted)
	for i := 0; i < len(values); i++ {
		if i < len(values)-1 && sorted[i] == sorted[i+1] {
			return false
		}
		if sorted[i] != values[i] {
			return false
		}
	}
	return true
}
