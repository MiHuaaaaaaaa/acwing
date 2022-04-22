package linkedlist

import (
	"fmt"
)

//带头单链表   有哨兵结点

type LinkedList struct {
	head   *ListNode
	length int
}

//结点
type ListNode struct {
	value interface{}
	next  *ListNode
}

func (l *ListNode) GetValue() interface{} {
	return l.value
}

func (l *ListNode) GetNext() *ListNode {
	return l.next
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{value: value, next: nil}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) (ok bool) {
	if nil == p {
		return
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) (ok bool) {
	if nil == p || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index int) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = cur.next
	p = nil
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += " -> "
		}
	}
	fmt.Println(format)
}

//单链表反转
func (this *LinkedList) Reverse() {
	if this.length == 0 || this.length == 1 {
		return
	}
	var pre *ListNode = nil
	cur := this.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	this.head.next = pre
}

//判断单链表是否有环
func (this *LinkedList) HasCycle() bool {
	if nil != this.head {
		slow := this.head
		fast := this.head
		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// MergeSortedList 合并两个有序链表
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	l := &LinkedList{head: &ListNode{}}
	cur := l.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) > cur2.value.(int) {
			cur.next = cur2
			cur2 = cur2.next
		} else {
			cur.next = cur1
			cur1 = cur1.next
		}
		cur = cur.next
	}
	if cur1 == nil {
		cur.next = cur2
	} else {
		cur.next = cur1
	}
	return l
}

// DeleteBottomN 删除倒数第N个节点
func (this *LinkedList) DeleteBottomN(n int) {
	slow, fast := this.head, this.head

	for i := 0; i < n && fast != nil; i++ {
		fast = fast.next
	}
	if nil == fast {
		return
	}

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next
}

//链表的中间结点
func (this *LinkedList) FindMiddleNode() *ListNode {
	slow := this.head
	fast := this.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
