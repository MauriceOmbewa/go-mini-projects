package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 2011, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	} else {
		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				if a[i] > a[j] {
					a[i], a[j] = a[j], a[i]
				}
			}
		}
		return a[len(a)-1]
	}
}
