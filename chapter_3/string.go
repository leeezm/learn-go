package main

import "fmt"

func main() {
	s := "left foot"
	t := s
	s += ",right foot"
	//s[0] = 'L'	字符串不可变
	fmt.Println(s)
	fmt.Println(t)
}
