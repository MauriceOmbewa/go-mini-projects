package main

import (
	"fmt"
	"os"
)

func main() {
	first := os.Args[1]
	second := os.Args[2]

	for _, char1 := range first {
		for _, char2 := range second {
			if char1 == char2 {
				fmt.Print(string(char1))
				break
			}
		}
	}
	fmt.Println()
}
