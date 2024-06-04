package main

import "fmt"

func main(){
	s := -123
	g := itoa(s)
	fmt.Println(g)
}

func itoa(s int) string{
	if s < 0{
		fmt.Print("-")
		s = s * -1
	}
	j := ""
	for s > 0{
		h := s%10
		j= string(h+'0')+j
		s /= 10
	}
	return j
}