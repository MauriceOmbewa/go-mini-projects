package main

import (
	"fmt"
	"os"
)
 func main(){
	if len(os.Args) != 2{
		return
	}
	input := os.Args[1]

	var result []rune

	for i, char := range input{
		if char >= 'a' && char <= 'z' && char + 13 > 'z'{
			result = append(result, rune(char + 13 - 26))
		} else if char >= 'A' && char <= 'Z' && char + 13 > 'Z'{
			result = append(result, rune(char + 13 - 26))
		} else if char >= 'a' && char <= 'z' && char + 13 < 'z'{
			result = append(result, rune(char + 13))
		}  else if char >= 'A' && char <= 'Z' && char + 13 < 'Z'{
			result = append(result, rune(char + 13))
		} else {
			result = append(result, rune(input[i]))
		}
	}
	fmt.Println(string(result))	
 }