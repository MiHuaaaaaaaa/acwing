package stack

import (
	"fmt"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	as := NewArrayStack()
	as.Push("1")
	as.Push("2")
	as.Push("3")
	as.Print()
}

func TestArrayStack_Pop(t *testing.T) {
	as := NewArrayStack()
	as.Push("1")
	as.Push("2")
	as.Push("3")
	as.Print()

	fmt.Println(as.Pop())
	fmt.Println(as.Pop())
	fmt.Println(as.Pop())

}

func TestArrayStack_Top(t *testing.T) {
	as := NewArrayStack()
	as.Push("1")
	as.Push("2")
	as.Push("3")

	fmt.Println(as.Top())
}

func TestArrayStack_Flush(t *testing.T) {
	as := NewArrayStack()
	as.Push("1")
	as.Push("2")
	as.Push("3")
	as.Print()

	as.Flush()
	as.Print()
}
