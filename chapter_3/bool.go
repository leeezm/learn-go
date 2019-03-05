//bool值练习
package main

import "fmt"

func main() {
	var a int
	a = 1
	// && 比 || 的优先级更高（&&：逻辑乘法，||：逻辑加法）
	if a > 10 || a > 0 && a < 3 {
		fmt.Println("success")
	}
}
