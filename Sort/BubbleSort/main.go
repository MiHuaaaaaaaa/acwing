package main

import "fmt"

func main() {
	nums := []int{2, 1, 643, 12, 52, 21}
	bubbleSort(nums)
	fmt.Println(nums)
}

func bubbleSort(q []int) {
	n := len(q)
	for i := 0; i <= n-1; i++ {
		for j := i; j <= n-1; j++ {
			if q[i] > q[j] {
				q[i], q[j] = q[j], q[i]
			}
		}
	}
}
