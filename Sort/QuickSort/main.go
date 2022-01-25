package main

import "fmt"

func main() {
	var arr = []int{3, 12, 5, 2, 12, 1, 63, 76}
	for _, v := range arr {
		fmt.Print(v, "\t")
	}

	fmt.Println()
	quickSort(arr, 0, len(arr)-1)
	fmt.Println()
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
}

/*
1. 确定分界点： **q[l]  ，  q[r]  ,   q[(l+r)/2] ，随机** 。随便挑一个值为**x**
2. **调整区间，左边区间所有的数 ≤x , 右边区间所有的数 ≥x（🌟）**
3. 递归处理左右两段
// */
func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	} //确定边界
	// 减1，加1是为了从让指针先移动再做比较
	i := l - 1         //左指针
	j := r + 1         //右指针
	x := arr[(l+r)>>1] //随便去一个值
	for i < j {
		for {
			i++ //左指针向右移动后在做比较
			if arr[i] >= x {
				break //左指针指向的值不满足<=x
			}
		}
		for {
			j-- //右指针向左移动后在做比较
			if arr[j] <= x {
				break //右指针指向的值不满足>=x
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	quickSort(arr, l, j)
	quickSort(arr, j+1, r)
}
