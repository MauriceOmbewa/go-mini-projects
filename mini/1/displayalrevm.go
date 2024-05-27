package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 0; i < 26; i++ {
		if i%2 != 0 {
			z01.PrintRune(rune('Z' - i))
		} else {
			z01.PrintRune(rune('z' - i))
		}
	}
	z01.PrintRune('\n')
}
