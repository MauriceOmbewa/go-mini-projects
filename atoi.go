package main

import "fmt"

func main(){
	s := "-123h"
	c := atoi(s)
	fmt.Println(c)
}

func atoi(s string) int{
	sign := 1
	j := 0

	if string(s[0]) == "-"{
		sign = -1
	}
	for _, char := range s[1:]{
		j = j*10+int(char)-'0'
	}
	return j*sign
}