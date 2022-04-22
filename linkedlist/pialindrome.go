package linkedlist

// 判断是否为回文单链表
func isPalindrome(l *LinkedList) bool {
	var vals []interface{}

	// 将链表存储到数组中
	cur := l.head.next
	for cur != nil {
		vals = append(vals, cur.value)
		cur = cur.next
	}

	// 双指针比较
	n := len(vals)
	for k, v := range vals[:n/2] {
		if v != vals[n-k-1] {
			return false
		}
	}
	return true
}
