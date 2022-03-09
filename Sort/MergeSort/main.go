package main

import "fmt"

var tmp = make([]int, 100010)

func main() {
	var q = []int{3, 12, 5, 2, 12, 1, 63, 76}
	for _, v := range q {
		fmt.Print(v, " ")
	}

	mergeSort(q, 0, len(q)-1)
	fmt.Println()
	for _, v := range q {
		fmt.Printf("%d ", v)
	}
}

func mergeSort(q []int, l, r int) {
	//1.确定分界点
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	//2.递归左右区间
	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)
	//3.定义i,j双指针,分别指向左右区间的开始,k 临时数组存储了多少个值
	k, i, j := 0, l, mid+1
	//4.i,j指针做比较，小的值添加到临时数组中，并向前移动，直到其中一指针走向尾部

	for i <= mid && j <= r {
		if q[i] <= q[j] {
			tmp[k] = q[i]
			k++
			i++
		} else {
			tmp[k] = q[j]
			k++
			j++
		}
	}
	//5.将i，j指针未走完的数值，添加到临时数组中
	for i <= mid {
		tmp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = q[j]
		k++
		j++
	}
	//6.将临时数组重新添加到原数组中
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = tmp[j]
	}
}
