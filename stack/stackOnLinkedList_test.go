package stack

import "testing"

func TestLinkedListStack_Push(t *testing.T) {
	stack := NewLinkedListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print()
}

func TestLinkedListStack_Pop(t *testing.T) {
	stack := NewLinkedListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())

	stack.Print()

}

func TestLinkedListStack_Top(t *testing.T) {
	stack := NewLinkedListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	t.Log(stack.Top())
}

func TestLinkedListStack_Flush(t *testing.T) {
	stack := NewLinkedListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print()

	stack.Flush()

	stack.Print()
}
