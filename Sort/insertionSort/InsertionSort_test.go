package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	q := []int{1, 263, 11, 12, 5, 12, 66, 2}
	InsertionSort(q)
	t.Log(q)
}
