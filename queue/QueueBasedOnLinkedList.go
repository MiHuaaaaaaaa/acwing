package queue

import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (this *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	if this.tail == nil {
		this.tail = node
		this.head = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

func (this *LinkedListQueue) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	res := this.head.val
	this.head = this.head.next
	this.length--
	return res
}

func (this *LinkedListQueue) String() string {
	if this.head == nil {
		return "empty queue"
	}
	res := "head"
	cur := this.head
	for cur != nil {
		res += fmt.Sprintf("<-%+v", cur.val)
		cur = cur.next
	}
	res += "<-tail"
	return res
}
