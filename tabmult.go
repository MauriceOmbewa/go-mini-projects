package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	inp := os.Args[1]

	inpconv, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ans int

	for i := 1; i <= 9; i++{
		ans = i*inpconv
		fmt.Println(i," x ","= ",ans)
	}
}
