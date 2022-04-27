package queue

import "testing"

func TestCircularQueue_EnQueue(t *testing.T) {
	queue := NewCircularQueue(3)
	t.Log(queue.EnQueue(1))
	t.Log(queue.EnQueue(2))
	t.Log(queue.EnQueue(3))
	t.Log(queue.EnQueue(4))

	t.Log(queue.String())
}

func TestCircularQueue_DeQueue(t *testing.T) {
	queue := NewCircularQueue(3)
	t.Log(queue.EnQueue(1))
	t.Log(queue.EnQueue(2))

	t.Log(queue.String())

	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())

	t.Log(queue.String())

}
