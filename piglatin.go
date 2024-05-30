package main

import (
	"os"
	"fmt"
)

func main(){
	if len(os.Args) != 2{
		return
	}

	input := os.Args[1]
	final := ""

	for i, char := range input{
		if (char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || 
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U') && (i == 0 ) {
			final = input + "ay"
		} else if (char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || 
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U') && (i != 0 ) {
			final = input[i:]
			for j := 0; j < i; j++{
				final += string(input[j])
			}
			final += "ay"
		}
	}
	fmt.Println(final)
}