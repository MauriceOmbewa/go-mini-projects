package main

import (
	"os"
	"fmt"
)

func main(){
	if len(os.Args) != 3{
		fmt.Println()
		return
	}
	first := os.Args[1]
	second := os.Args[2]

	final := ""

	for _, r := range first{
		status := true
		for _, g := range final{
			// fmt.Println(g)
			// fmt.Println(r)
			if g == r{
				status = false
				break
			}
		}
		if status{
			final = final + string(r)
			continue
		}
		
	}

	for _, r := range second{
		status := true
		for _, g := range final{
			if g == r{
				status = false
				break
			}
		}
		if status{
			final = final + string(r)
		}
		
	}
	fmt.Println(final)
	fmt.Println()
}