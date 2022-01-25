package main

import "fmt"

// 题意：升序整数数组，查找元素k的起始位置和终止位置（位置从 0 开始计数）
func main() {
	// var n, k int
	// fmt.Scanf("%d%d", &n, &k)
	// var q = make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%d", &q[i])
	// }
	k := 3
	q := []int{1, 2, 2, 3, 3, 4}
	x := []int{5, 4, 3}
	for ; k > 0; k-- {
		fmt.Println(x[k-1])
		l := bsearch_1(q, x[k-1])
		r := bsearch_2(q, x[k-1])
		fmt.Printf("%d %d\n", l, r)
	}
}

func bsearch_1(q []int, x int) int {
	l, r := 0, len(q)-1
	for l < r {
		mid := (l + r) >> 1
		if q[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if q[l] != x {
		return -1
	}
	return l
}

func bsearch_2(q []int, x int) int {
	l, r := 0, len(q)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if q[mid] <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if q[l] != x {
		return -1
	}
	return l
}
