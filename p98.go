package main

func isValidBST(root *TreeNode) bool {
	var fx func(*TreeNode) bool
	var prevNode *TreeNode
	fx = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if fx(root.Left) == false {
			return false
		}
		if prevNode != nil && prevNode.Val >= root.Val {
			return false
		}
		prevNode = root
		return fx(root.Right)
	}
	return fx(root)
}
