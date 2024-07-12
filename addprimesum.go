package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
		return
	}
	input := os.Args[1]
	inpconv, _ := strconv.Atoi(input)
	sum := 0

	for i := inpconv; i >=1; i--{
		if prime(i){
			sum += i
		} 
	}
	fmt.Println(sum)
}

func prime(a int) bool{
	if a == 1{
		return false
	} else if a == 2{
		return true
	}
	for i := 2; i < a; i++{
		if a%i == 0{
			return false
		} 
	}
	return true
}
