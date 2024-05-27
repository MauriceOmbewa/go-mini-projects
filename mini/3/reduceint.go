package main

import (
	"fmt"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}

	acc := a[0]
	// cur := a[1]
	// acc = f(acc, cur)
	// fmt.Println(acc)

	for _, cur := range a[1:] {
		acc = f(acc, cur)
	}
	fmt.Println(acc)
}
