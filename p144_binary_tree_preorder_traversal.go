package main

func preorderDFS(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	*values = append(*values, root.Val)
	preorderDFS(root.Left, values)
	preorderDFS(root.Right, values)

}
func preorderTraversal(root *TreeNode) []int {
	values := []int{}
	preorderDFS(root, &values)
	return values

}
