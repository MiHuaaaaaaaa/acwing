package main

// Element 用于表示链表当中的节点
type Element struct {
	// next, prev 分别表示上个节点和下个节点
	next, prev *Element

	// 表示节点所在的元素
	list *List

	// 节点所保存的数据，也就是上面图中的 data
	Value interface{}
}

// List 这是一个双向链表
// List 的零值是一个空链表
type List struct {
	// 根节点，List 其实是一个双向循环链表，root， root.prev 是尾节点, 尾节点的下一个节点指向 root
	// 根节点是一个哨兵节点，是为了用来简化节点操作使用的
	root Element

	// 链表的长度，不包括哨兵节点，也就是根节点
	len int
}

// 将节点 e 插入 at 之后
func (l *List) insert(e, at *Element) *Element {
	// 假设 at.next 为 nt
	// 1. 将节点 e 的上一个节点指向 at
	e.prev = at
	// 2. 将节点 e 的下一个节点指向 nt
	e.next = at.next
	// 3. 这个时候  e.prev.next == at.next
	// 其实就是本来 at --> nt，修改为 at --> e
	e.prev.next = e
	// 4. e.next.prev == nt.prev
	// 本来 at <--- nt，修改为 e <--- nt
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// remove removes e from its list, decrements l.len, and returns e.
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	// 这里为了避免内存泄漏的操作可以学习
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}
