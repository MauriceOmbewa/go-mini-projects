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

	inpconv, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	final := strconv.FormatInt(inpconv, 16)
	fmt.Println(final)
}