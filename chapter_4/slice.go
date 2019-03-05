package main

import "fmt"

func main() {
	slice := []string{1: "January", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September",
		10: "Oct", 11: "Nov", 12: "Dec"}
	//超过cap()大小，导致程序出错
	//fmt.Println(cap(slice), slice[:14])
	summer := slice[6:9]
	fmt.Println(summer)
	//在slice范围内扩展了slice
	endlesssummer := summer[:7]
	fmt.Println(endlesssummer)
}
