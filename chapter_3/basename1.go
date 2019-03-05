package main

import (
	"fmt"
)

func main() {
	//eg a=>a, a.go=>a, a/b/c.go=>c, a/b.c.go=>b.c
	s := "a/b.c.go"
	//最后一个 / 前面的省略
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	//最后一个 . 后面的全部省略
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	fmt.Println(s)
}
