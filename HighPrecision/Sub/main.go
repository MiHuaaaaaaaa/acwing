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
		cSlice []int
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

	// 2.比较两个值的大小
	if cmp(aSlice, bSlice) { //a>b
		cSlice = sub(aSlice, bSlice)
		// 3. 打印结果
		for i := len(cSlice) - 1; i >= 0; i-- {
			fmt.Printf("%d", cSlice[i])
		}
	} else { //b>a
		cSlice = sub(bSlice, aSlice)
		// 3. 打印结果
		fmt.Print("-")
		for i := len(cSlice) - 1; i >= 0; i-- {
			fmt.Printf("%d", cSlice[i])
		}
	}

	fmt.Println()
}

// c = a - b
func sub(a, b []int) (c []int) {
	// t 进位
	//Ci=Ai-Bi-t
	for i, t := 0, 0; i < len(a); i++ {
		t = a[i] - t
		if i < len(b) {
			t -= b[i]
		}
		// 为什么 (t+10)%10
		// a-b = t
		// 如果 t> 0 结果为  t
		// t<0  结果为  t+10
		// 两者相结合  (t+10)%10
		// eg. 22-13 ==>   2 - 3 <0,向前借一位  ==> 12-3=9 相等于 （-1+10）%10=9
		c = append(c, (t+10)%10)
		if t < 0 { //证明向前借了一位
			t = 1
		} else {
			t = 0
		}
	}
	for len(c) > 1 && c[len(c)-1] == 0 { //去掉前导0
		c = c[:len(c)-1]
	}
	return
}

// 判断 a>=b
func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}
