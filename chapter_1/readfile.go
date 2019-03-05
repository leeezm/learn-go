package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("Absolute file path")
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
	}
	f.Close()
}
