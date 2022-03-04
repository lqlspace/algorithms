package simple

type Stack struct {
	items []int
}

func (s *Stack) Empty() bool {
	return s.Size() ==  0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() int {
	val := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]

	return val
}

func (s *Stack) Top() int {
	return s.items[s.Size()-1]
}


type MyQueue struct {
	in, out Stack
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	this.in.Push(x)
}


func (this *MyQueue) Pop() int {
	if !this.out.Empty() {
		return this.out.Pop()
	}

	for !this.in.Empty() {
		v := this.in.Pop()
		this.out.Push(v)
	}

	return this.out.Pop()
}


func (this *MyQueue) Peek() int {
	if !this.out.Empty() {
		return this.out.Top()
	}

	for !this.in.Empty() {
		v := this.in.Pop()
		this.out.Push(v)
	}

	return this.out.Top()
}


func (this *MyQueue) Empty() bool {
	return this.in.Empty() && this.out.Empty()
}
