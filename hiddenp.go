package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 3{
		return
	}
	first := os.Args[1]
	second := os.Args[2]

	index := 0
	if first == ""{
		fmt.Println("1")
		return
	}

	for _, char := range second{
		if char == rune(first[index]){
			index++
		}
		if index == len(first){
			break
		}
	}

	if index == len(first){
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}