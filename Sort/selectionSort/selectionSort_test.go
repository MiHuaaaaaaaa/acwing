package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	q := []int{1, 263, 11, 12, 5, 12, 66, 2}
	SelectionSort(q)
	t.Log(q)
}
