package main

func dfs1(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	dfs1(root.Left, values)
	dfs1(root.Right, values)
	*values = append(*values, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	values := []int{}

	dfs1(root, &values)
	return values
}
