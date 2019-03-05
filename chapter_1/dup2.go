package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		count(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			count(f, counts)
			f.Close()
			for _, num := range counts {
				if num > 1 {
					fmt.Printf("%s\n", file)
				}
			}
			counts = make(map[string]int)
		}
	}
}

func count(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
