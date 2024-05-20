package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for j := 0; j < 26; j++ {
		if j%2 == 0 {
			z01.PrintRune(rune('a' + j))
		} else {
			z01.PrintRune(rune('A' + j))
		}
	}
	z01.PrintRune('\n')
}
