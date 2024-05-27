package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	first := os.Args[1]
	second := os.Args[2]
	i := 0
	for _, char := range second {
		if i < len(first) && char == rune(first[i]) {
			i++
		}
		if i == len(first) {
			break
		}
	}
	if i == len(first) {
		for _, char := range first {
			z01.PrintRune(rune(char))
		}
		z01.PrintRune('\n')
	}
}
hello h3ell4odfgry