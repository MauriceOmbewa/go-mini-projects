package main

import (
	"fmt"
	"strings"
	"os"
)

func main(){
	if len(os.Args) != 2{
		return
	}

	var final []string

	input := os.Args[1]
	inp := strings.Fields(input)

	for i := len(inp)-1; i >= 0; i--{
		final = append(final, inp[i])
	}

	ans := strings.Join(final, " ")
	fmt.Println(ans)
}