package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	q := []int{1, 263, 11, 12, 5, 12, 66, 2}
	BubbleSort(q)
	t.Log(q)
}
