package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	input := os.Args[1]
	var final []rune

	for _, char := range input {
		if char >= 'a' && char <= 'z' && char+13 > 'z' {
			final = append(final, rune(char+13-26))
		} else if char >= 'A' && char <= 'Z' && char+13 > 'Z' {
			final = append(final, rune(char+13-26))
		} else if char >= 'a' && char <= 'z' && char+13 <= 'z' {
			final = append(final, rune(char+13))
		} else if char >= 'A' && char <= 'Z' && char+13 <= 'Z' {
			final = append(final, rune(char+13))
		} else {
			final = append(final, char)
		}
	}
	fmt.Println(string(final))
}
