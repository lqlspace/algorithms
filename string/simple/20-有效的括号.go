package simple

// 方法1：创建一个stack
func isValid1(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		e := s[i]
		if e == '(' || e == '{' || e == '[' {
			stack.Push(e)
		}
		if e == ')' || e == ']' || e == '}' {
			es := stack.Pop()
			if es == 0 || !match(es, e) {
				return false
			}
		}
	}
	if stack.Len() > 0 {
		return false
	}

	return true
}

type Stack struct {
	entries []byte
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(b byte) {
	s.entries = append(s.entries, b)
}

func (s *Stack) Pop() byte {
	l := len(s.entries)
	if l == 0 {
		return 0
	}
	res := s.entries[l-1]
	s.entries = s.entries[:l-1]

	return res
}

func (s *Stack) Len() int {
	return len(s.entries)
}


func match(s, d byte) bool {
	switch {
	case s == '(' && d == ')':
		return true
	case s == '[' && d == ']':
		return true
	case s == '{' && d == '}':
		return true
	default:
		return false
	}
}

//方法2
func isValid2(s string) bool {
	if len(s) % 2 == 1 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		e := getValue(s[i])
		if e == 0 {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || e != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func getValue(s byte) byte {
	switch s {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	default:
		return 0
	}
}

