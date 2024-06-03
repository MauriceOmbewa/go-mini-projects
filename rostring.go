package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	if len(os.Args) != 2{
		fmt.Println()
		return
	}

	input := os.Args[1]
	mdoinp := ""

	for i, char := range input{
		if char == ' ' && rune(input[i+1]) == ' '{
			continue
		} else {
			mdoinp += string(char)
		}
	}
	fin := strings.Fields(mdoinp)

	var finalstr string
	for i := 0; i < len(fin); i++{
		if i > 0{
			finalstr += fin[i] + " "
		}
	}
	finalstr += fin[0]
	fmt.Println(finalstr)
	
}