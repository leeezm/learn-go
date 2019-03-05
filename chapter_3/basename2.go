package main

import (
	"fmt"
	"strings"
)

func main() {
	//eg a=>a, a.go=>a, a/b/c.go=>c, a/b.c.go=>b.c
	s := "a"
	//最后一个 / 前面的省略
	index := strings.LastIndex(s, "/")
	s = s[index+1:]
	//最后一个 . 后面的全部省略
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	fmt.Println(s)
}
