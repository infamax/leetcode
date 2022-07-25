type Queue struct {
	data      []*Node
	firstItem int
}

func New() *Queue {
	return &Queue{
		data:      []*Node{},
		firstItem: 0,
	}
}

func (q *Queue) Push(x *Node) {
	q.data = append(q.data, x)
}

func (q *Queue) Top() *Node {
	return q.data[q.firstItem]
}

func (q *Queue) Pop() {
	q.firstItem++
}

func (q *Queue) Length() int {
	return len(q.data) - q.firstItem
}

func max(depths map[*Node]int) int {
	res := 0
	for _, value := range depths {
		if value > res {
			res = value
		}
	}
	return res
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depths := make(map[*Node]int)
	depths[root] = 1
	q := New()
	q.Push(root)
	for q.Length() > 0 {
		v := q.Top()
		for _, child := range v.Children {
			depths[child] = depths[v] + 1
			q.Push(child)
		}
		q.Pop()
	}
	return max(depths)
}
