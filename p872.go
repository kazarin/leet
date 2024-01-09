package main

func dfs872(root *TreeNode, acc *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {

		*acc = append(*acc, root.Val)
	}
	dfs872(root.Left, acc)
	dfs872(root.Right, acc)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1, leafs2 := []int{}, []int{}
	dfs872(root1, &leafs1)
	dfs872(root2, &leafs2)
	if len(leafs1) != len(leafs2) {
		return false
	}
	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}
	return true
}
