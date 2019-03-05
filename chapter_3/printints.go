package main

import (
	"bytes"
	"fmt"
)

func intToString(values []int) string {
	var buf bytes.Buffer
	//添加任意文字符号的UTF-8编码，最好使用 WriteRune 方法
	buf.WriteByte('[')
	for index, val := range values {
		if index > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", val)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intToString([]int{1, 2, 3, 4, 5})) //[1,2,3]
}
