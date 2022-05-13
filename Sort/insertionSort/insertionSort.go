package sort

func InsertionSort(q []int) {
	n := len(q)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && q[j] < q[j-1]; j-- {
			q[j], q[j-1] = q[j-1], q[j]
		}
	}
}
