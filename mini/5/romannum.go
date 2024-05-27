package main

import (
	"os"
	"fmt"
	"strconv"
)

func main(){
	if len(os.Args) != 2{
		return
	}

	input := os.Args[1]

	inputconv, _ := strconv.Atoi(input)

	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	numbers := []
}