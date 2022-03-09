package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	var (
		aSlice = make([]int, 0)
		bSlice = make([]int, 0)
	)
	// 1.大整数存储到数组
	// 倒着存储，低位表示个数，高位表示大数
	// eg. "123456"   ==>   {[6],[5],[4],[3],[2],[1]}
	for i := len(a) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(string(a[i]))
		aSlice = append(aSlice, v)
	}

	for i := len(b) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(string(b[i]))
		bSlice = append(bSlice, v)
	}

	// 2.相加  模拟人工相加
	cSlice := add(aSlice, bSlice)
	// 3. 打印结果
	for i := len(cSlice) - 1; i >= 0; i-- {
		fmt.Printf("%d", cSlice[i])
	}
	fmt.Println()
}

// 每一位数为 Ai+Bi+t
func add(a, b []int) (c []int) {
	t := 0 //进位
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) {
			t += a[i]
		}
		if i < len(b) {
			t += b[i]
		}
		c = append(c, t%10)
		t /= 10
	}
	if t > 0 {
		c = append(c, t)
	}
	return
}
