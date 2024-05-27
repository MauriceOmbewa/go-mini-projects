package main

import (
	"fmt"
)

func main() {
	n := 66
	var boool bool
	for n/2 != 1 {
        if n%2 == 0 {
            boool = true
        } else {
			boool = false
		}
        n = n / 2
    }
	fmt.Println(boool)
}
// if n%2 != 0 {
// 	fmt.Println("False")
// } else 