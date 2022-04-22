package linkedlist

import "testing"

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}

func TestReverse(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	l.Reverse()
	l.Print()

}

func TestMergeSortedList(t *testing.T) {
	l1 := NewLinkedList()
	for i := 0; i < 3; i++ {
		l1.InsertToTail(i + 1)
	}

	l2 := NewLinkedList()
	for i := 0; i < 5; i++ {
		l2.InsertToTail(i + 1)
	}
	l := MergeSortedList(l1, l2)
	l.Print()
}

func TestDeleteBottomN(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	l.DeleteBottomN(2)
	l.Print()
}

func TestFindMiddleNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 6; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	node := l.FindMiddleNode()
	t.Log(node.value)
}
