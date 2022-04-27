package queue

import "testing"

func TestLinkedList_EnQueue(t *testing.T) {
	queue := NewLinkedListQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)

	t.Log(queue.String())
}

func TestLinkedList_DeQueue(t *testing.T) {
	queue := NewLinkedListQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)

	t.Log(queue.String())

	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())

	t.Log(queue.String())
}
