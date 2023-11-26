package main

func dfs(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0

	}
	return dfs(root.Left, targetSum) || dfs(root.Right, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	answer := dfs(root, targetSum)

	return answer

}
