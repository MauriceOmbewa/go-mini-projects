package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]

	for i, char := range input {
		if char == ' ' {
			for j := i + 1; j < len(input); j++ {
				z01.PrintRune(rune(input[j]))
			}
		}
	}
	z01.PrintRune('\n')
}
