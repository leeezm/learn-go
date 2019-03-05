package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, path := range os.Args[1:] {
		file, _ := ioutil.ReadFile(path)
		for _, line := range strings.Split(string(file), "\n") {
			counts[line]++
		}
	}
	for index, num := range counts {
		if num > 1 {
			fmt.Printf("%s\t%d\n", index, num)
		}
	}
}
