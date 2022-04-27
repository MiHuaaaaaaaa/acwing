package queue

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (this *ArrayQueue) EnQueue(v interface{}) bool {
	if this.capacity == this.tail {
		return false
	}
	this.q[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) DeQueue() interface{} {
	if this.head == this.tail {
		return nil
	}
	res := this.q[this.head]
	this.head++
	return res
}

func (this *ArrayQueue) String() string {
	if this.head == this.tail {
		return "empty queue"
	}
	result := "head"
	for _, v := range this.q {
		result += fmt.Sprintf("<-%+v", v)
	}
	result += "<-tail"
	return result
}
