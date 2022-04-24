package stack

import "fmt"

type LinkedListStack struct {
	topNode *node
}

var _ = Stack(&LinkedListStack{})

type node struct {
	next *node
	val  interface{}
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

// Push 入栈操作
func (s *LinkedListStack) Push(v interface{}) {
	node := &node{val: v}
	node.next = s.topNode
	s.topNode = node
}

// Pop 出栈操作
func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	pop := s.topNode.val
	s.topNode = s.topNode.next
	return pop
}

func (s *LinkedListStack) Top() interface{} {
	return s.topNode.val
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.topNode == nil
}

func (s *LinkedListStack) Flush() {
	s.topNode = nil
}

func (s *LinkedListStack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty stack")
		return
	}
	cur := s.topNode
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}
