package main

import "fmt"

func main() {
	//switch getRes() {
	//case "aaaa":
	//	fmt.Println("aaaa")
	//case "bbbb":
	//	fmt.Println("bbbb")
	//default:
	//	fmt.Println("cccc")
	//}

	x := -1
	switch {
	case x > 0:
		fmt.Println("大于0")
	case x < 0:
		fmt.Println("小于0")
	default:
		fmt.Println("等于0")
	}
}

func getRes() string {
	return "aaaa"
}
