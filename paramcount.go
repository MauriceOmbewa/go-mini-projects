package main

import (
	"fmt"
	"os"
)

func main() {
	var count int
	for i := 1; i < len(os.Args); i++ {
		count++
	}
	fmt.Println(count)
}
