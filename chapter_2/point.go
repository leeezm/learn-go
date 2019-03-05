package main

import "fmt"

func main() {
	//通过调用test产生的局部变量v在函数调用返回后依然存在
	p := test()
	fmt.Println(*p)
	*p = 3
	fmt.Println(*p)
}

func test() *int {
	v := 1
	return &v
}
