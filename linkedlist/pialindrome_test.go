package linkedlist

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}

	for i := 3; i > 0; i-- {
		l.InsertToTail(i)
	}

	l.Print()
	// fmt.Println(isPalindrome(l))
}
