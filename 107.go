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

func (q *Queue) Pop() {
	q.firstItem++
}

func (q *Queue) Top() *TreeNode {
	return q.data[q.firstItem]
}

func (q *Queue) Length() int {
	return len(q.data) - q.firstItem
}

func reverse(data [][]int) {
	for i := 0; i < len(data) / 2; i++ {
		data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	levels := make([][]int, 0)
	levels = append(levels, []int{root.Val})
	level := make([]int, 0)
	q := New()
	q.Push(root)
	q.Push(nil)
	for q.Length() > 0 {
		v := q.Top()

		if v == nil {
			if len(level) != 0 {
				levels = append(levels, level)
			}
			level = make([]int, 0)
			if q.Length() != 1 {
				q.Push(nil)
			}
			q.Pop()
			continue
		}

		if v.Left != nil {
			q.Push(v.Left)
			level = append(level, v.Left.Val)
		}

		if v.Right != nil {
			q.Push(v.Right)
			level = append(level, v.Right.Val)
		}

		q.Pop()
	}
	reverse(levels)
	return levels
}
