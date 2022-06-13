type Stack struct {
	data []int
	n    int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
		n:    0,
	}
}

func (s *Stack) Push(x int) {
	s.n++
	s.data = append(s.data, x)
}

func (s *Stack) Pop() {
	s.n--
	s.data = s.data[:s.n]
}

func (s *Stack) Peek() int {
	return s.data[s.n-1]
}

func (s *Stack) Empty() bool {
	return s.n == 0
}

type MyQueue struct {
	s1 *Stack
	s2 *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (q *MyQueue) Push(x int) {
	q.s1.Push(x)
}

func (q *MyQueue) Pop() int {
	for !q.s1.Empty() {
		q.s2.Push(q.s1.Peek())
		q.s1.Pop()
	}

	res := q.s2.Peek()
	q.s2.Pop()	

	for !q.s2.Empty() {
		q.s1.Push(q.s2.Peek())
		q.s2.Pop()
	}
	
	return res
}

func (q *MyQueue) Peek() int {
	for !q.s1.Empty() {
		q.s2.Push(q.s1.Peek())
		q.s1.Pop()
	}

	res := q.s2.Peek()

	for !q.s2.Empty() {
		q.s1.Push(q.s2.Peek())
		q.s2.Pop()
	}

	return res
}

func (q *MyQueue) Empty() bool {
	return q.s1.Empty()
}
