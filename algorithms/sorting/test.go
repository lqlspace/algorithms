/*
* 将int型变量a的第k位清0，即a = a & ^(1 << k)
* 将int型变量a的第k位置1，即a = a | (1 << k)
*/
package main

import (
	"fmt"
	"io"
)

const (
	SHIFT  = 5
	MASK  = 31

)

//相应位置1
func set(a []int, i int) {
	a[i>>SHIFT] |= (1<<uint((i & MASK)))  //等价于a[i/32] |= (1 << uint(i % 32))
}


//相应位置0
func clr(a []int, i int) {
	a[i>>SHIFT] &= ^(1<<uint((i & MASK)))
}

func test(a []int, i int) bool {
	if a[i>>SHIFT] & (1 << uint((i & MASK))) > 0 {
		return true
	}
	return false
}


func main() {
	items := make([]int, 100)
	for i := 0; i < 100; i++ {
		clr(items, i)
	}
	var tmp int
	for {
		_, err := fmt.Scanf("%d", &tmp)
		if err == io.EOF {
			break
		}
		set(items, tmp)
	}

	fmt.Printf("\n")
	for i := 0; i < 100; i++ {
		if test(items, i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}

