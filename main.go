package main

import "fmt"

var tmp = make([]int, 100001)

func mergeSort(q []int, l, r int, cnt *int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, l, mid, cnt)
	mergeSort(q, mid+1, r, cnt)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			tmp[k] = q[i]
			k++
			i++
		} else {
			tmp[k] = q[j]
			// 				q[i]>q[j] 且 i<j 满足逆序对
			//  l 		____i____________ mid
			//  mid+1	________j________  r
			//当第i个位置的值大于，j位置的值时，i之后的所有值全部都大于j，也就都是逆序对的数量= mid - i + 1。
			*cnt += mid - i + 1
			k++
			j++
		}
	}
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
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = tmp[j]
	}
}

func main() {
	// var n int
	// fmt.Scanf("%d", &n)

	// nums := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%d", &nums[i])
	// }
	var nums = []int{3, 12, 5, 2, 12, 1, 63, 76}
	cnt := 0
	mergeSort(nums, 0, len(nums)-1, &cnt)
	fmt.Println(cnt)
	for _, v := range nums {
		fmt.Print(v, " ")
	}
}
