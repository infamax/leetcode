func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}

	depthLeftTree := 1 + minDepth(root.Left)
	depthRightTree := 1 + minDepth(root.Right)
	return min(depthLeftTree, depthRightTree)
}
