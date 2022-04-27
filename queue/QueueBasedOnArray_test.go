package queue

import "testing"

func TestEnQueue(t *testing.T) {
	queue := NewArrayQueue(3)
	t.Log(queue.EnQueue(1))
	t.Log(queue.EnQueue(2))
	t.Log(queue.EnQueue(3))
	t.Log(queue.EnQueue(4))
	t.Log(queue.String())
}

func TestDeQueue(t *testing.T) {
	queue := NewArrayQueue(3)
	t.Log(queue.EnQueue(1))
	t.Log(queue.EnQueue(2))
	t.Log(queue.EnQueue(3))

	t.Log(queue.String())

	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())

	t.Log(queue.String())
}
