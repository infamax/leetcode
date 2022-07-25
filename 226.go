type Queue struct {
	data      []*TreeNode
	firstItem int
}

func New() *Queue {
	return &Queue{
		data:      []*TreeNode{},
		firstItem: 0,
	}
}

func (q *Queue) Push(x *TreeNode) {
	q.data = append(q.data, x)
}

func (q *Queue) Top() *TreeNode {
	return q.data[q.firstItem]
}

func (q *Queue) Pop() {
	q.firstItem++
}

func (q *Queue) Length() int {
	return len(q.data) - q.firstItem
}

func dfs(root *TreeNode) {
	q := New()
	q.Push(root)
	for q.Length() > 0 {
		v := q.Top()
		v.Left, v.Right = v.Right, v.Left
		if v.Left != nil && v.Right != nil {
			q.Push(v.Right)
			q.Push(v.Left)
		} else if v.Left != nil {
			q.Push(v.Left)
		} else if v.Right != nil {
			q.Push(v.Right)
		}
		q.Pop()
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dfs(root)
	return root
}
