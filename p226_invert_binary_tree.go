package main

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	traverse(root.Left)
	traverse(root.Right)

}
func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}
