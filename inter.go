package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	// Check if the number of arguments is 2
	if len(os.Args) != 3 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]

	index := 0
	// Slice to store characters found in both strings
	var commonChars []rune

	// Loop through each character in the second string
	for _, char := range second {
		// If the character matches the character in the first string at the current index
		if rune(first[index]) == char {
			// Increment the index to check the next character in the first string
			index++
			// Store the character in commonChars if it's not already there
			found := false
			for _, c := range commonChars {
				if c == char {
					found = true
					break
				}
			}
			if !found {
				commonChars = append(commonChars, char)
			}
		}
		// If all characters from the first string are found in the second string
		if len(first) == index {
			break
		}
	}

	// Print the characters found in both strings
	for _, char := range commonChars {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
