//堆栈内存分配和变量逃逸
package main

var global *int

func f() {
	var x int
	//堆空间分配，x从f中逃逸
	x = 1
	global = &x
}

func g() {
	//栈分配
	y := new(int)
	*y = 1
}
