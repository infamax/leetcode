func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter func(root *TreeNode) int
	maxD := 0
	diameter = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftD := diameter(root.Left)
		rightD := diameter(root.Right)
		maxD = max(maxD, leftD+rightD)
		return 1 + max(leftD, rightD)
	}
	diameter(root)
	return maxD
}
