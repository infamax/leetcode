func preorder(root *Node) []int {
	res := make([]int, 0)

	var preorderTree func(root *Node)

	preorderTree = func(root *Node) {
		if root == nil {
			return
		}
        res = append(res, root.Val)
		for _, v := range root.Children {
			preorderTree(v)
		}
	}

	preorderTree(root)

	return res
}
