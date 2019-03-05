package main

import "fmt"

const (
	One int = iota
	Two
	Three
	Four
	Five
)

func main() {
	fmt.Println(One, Two, Three, Four, Five)
}
