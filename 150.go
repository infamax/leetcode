type Stack struct {
	data []int
}

func New() *Stack {
	return &Stack{
		data: []int{},
	}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func evaluate(op1, op2 int, operation string) int {
	if operation == "+" {
		return op1 + op2
	} else if operation == "-" {
		return op1 - op2
	} else if operation == "*" {
		return op1 * op2
	}
	return op1 / op2
}

var operations = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
	"/": {},
}

func evalRPN(tokens []string) int {
	s := New()
	for _, token := range tokens {
		_, ok := operations[token]
		if ok {
			op2 := s.Top()
			s.Pop()
			op1 := s.Top()
			s.Pop()
			s.Push(evaluate(op1, op2, token))
		} else {
			num, _ := strconv.Atoi(token)
			s.Push(num)
		}
	}
	return s.Top()
}
