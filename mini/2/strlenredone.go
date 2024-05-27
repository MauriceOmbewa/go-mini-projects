package main

import (
	"fmt"
)

func main() {
	l := StrLen("ÈHello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	var count int
	for range s {
		count++
	}
	return count
}
