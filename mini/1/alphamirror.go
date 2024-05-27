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
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			z01.PrintRune('z' - (char - 'a'))
		} else if char >= 'A' && char <= 'Z' {
			z01.PrintRune('Z' - (char - 'A'))
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
