package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]
	var final []int

	fconv, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println(err)
		return
	}

	sconv, err := strconv.Atoi(second)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 2; i < fconv || i < sconv; i++ {
		if fconv%i == 0 && sconv%i == 0 {
			final = append(final, i)
			fconv /= i
			sconv /= i
		}else if fconv%i != 0 && sconv%1 != 0 {
			break
		}
	}
	tfinal := 1

	// ans := 0
	for i := 0; i < len(final); i++{
		// fmt.Println(tfinal)
		tfinal*= final[i]
		// fmt.Println(tfinal)
	}
	fmt.Println(tfinal)
}
