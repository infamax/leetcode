type StackString struct {
	data []string
}

func NewStringStack() *StackString {
	return &StackString{
		data: []string{},
	}
}

func (s *StackString) Push(x string) {
	s.data = append(s.data, x)
}

func (s *StackString) Top() string {
	return s.data[len(s.data)-1]
}

func (s *StackString) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *StackString) Empty() bool {
	return len(s.data) == 0
}

type StackInt struct {
	data []int
}

func NewIntStack() *StackInt {
	return &StackInt{
		data: []int{},
	}
}

func (s *StackInt) Push(x int) {
	s.data = append(s.data, x)
}

func (s *StackInt) Top() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *StackInt) Pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *StackInt) Empty() bool {
	return len(s.data) == 0
}

var priorityOperations = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

var validOperations = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
	"/": {},
	"(": {},
	")": {},
}

var operations = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
	"/": {},
}

func Tokens(s string) []string {
	tokens := make([]string, 0)
	var numStr strings.Builder
	for _, ch := range s {
		if ch == ' ' {
			continue
		}
		_, ok := validOperations[string(ch)]
		if ok {
			if numStr.Len() != 0 {
				tokens = append(tokens, numStr.String())
				numStr.Reset()
			}
			tokens = append(tokens, string(ch))
		} else {
			numStr.WriteRune(ch)
		}
	}
	if numStr.Len() != 0 {
		tokens = append(tokens, numStr.String())
		numStr.Reset()
	}
	return tokens
}

func GetReversePolishNotation(tokens []string) []string {
	reversePolishNotation := make([]string, 0)
	s := NewStringStack()
	for _, token := range tokens {
		_, err := strconv.Atoi(token)
		if err == nil {
			reversePolishNotation = append(reversePolishNotation, token)
		} else {
			if token == "(" {
				s.Push(token)
			} else if token == ")" {
				for s.Top() != "(" {
					reversePolishNotation = append(reversePolishNotation, s.Top())
					s.Pop()
				}
				s.Pop()
			} else {
				for !s.Empty() && priorityOperations[s.Top()] >= priorityOperations[token] {
					reversePolishNotation = append(reversePolishNotation, s.Top())
					s.Pop()
				}
				s.Push(token)
			}
		}
	}

	for !s.Empty() {
		reversePolishNotation = append(reversePolishNotation, s.Top())
		s.Pop()
	}

	return reversePolishNotation
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

func evalReversePolishNotations(tokens []string) int {
	s := NewIntStack()
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

func calculate(s string) int {
	tokens := Tokens(s)
	tokens = GetReversePolishNotation(tokens)
	return evalReversePolishNotations(tokens)
}
