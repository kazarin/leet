package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTree(array []interface{}) *TreeNode {
	var fx func(int) *TreeNode
	fx = func(i int) *TreeNode {
		if i >= len(array) || array[i] == nil {
			return nil
		}

		node := &TreeNode{Val: array[i].(int)}
		node.Left = fx(2*i + 1)
		node.Right = fx(2*i + 2)
		return node
	}
	return fx(0)
}
