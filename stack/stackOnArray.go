package stack

import "fmt"

type ArrayStack struct {
	items []interface{} //
	top   int           //栈顶指针
}

var _ = Stack(&ArrayStack{})

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		items: make([]interface{}, 0, 32),
		top:   -1,
	}
}

// Push 入栈操作
func (s *ArrayStack) Push(item interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top++
	}
	if condition := s.top < len(s.items); condition {
		s.items[s.top] = item
	} else {
		s.items = append(s.items, item)
	}
}

// Pop 出栈操作
func (s *ArrayStack) Pop() interface{} {
	item := s.items[s.top]
	s.top--
	return item
}

func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return "empty statck"
	}
	return s.items[s.top]
}

func (s *ArrayStack) IsEmpty() bool {
	return s.top < 0
}

func (s *ArrayStack) Flush() {
	s.top = -1
}

func (s *ArrayStack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty statck")
		return
	}
	for _, v := range s.items {
		fmt.Println(v)
	}
}
