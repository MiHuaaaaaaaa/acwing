package sort

func BubbleSort(q []int) {
	n := len(q)
	for i := 0; i < n-1; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if q[j] > q[j+1] {
				q[j], q[j+1] = q[j+1], q[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
