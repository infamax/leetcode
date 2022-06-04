type MinStack struct {
	data []int
	min  int
	n    int
}

const inf = math.MaxInt - 1

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  inf,
		n:    0,
	}
}

func (m *MinStack) Push(val int) {
	m.min = min(m.min, val) // update min
	m.n++
	m.data = append(m.data, val)
}

func (m *MinStack) Pop() {
	m.n--
	m.data = m.data[:m.n]
	m.min = updateMin(m.data)
}

func (m *MinStack) Top() int {
	return m.data[m.n-1]
}

func (m *MinStack) GetMin() int {
	return m.min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func updateMin(a []int) int {
	if len(a) == 0 {
		return inf
	}
	minItem := a[0]
	for i := 1; i < len(a); i++ {
		minItem = min(minItem, a[i])
	}
	return minItem
}
