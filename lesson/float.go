package main

import (
	"fmt"
)

func main() {
	var box float64
	for count := 0; count < 11; count++ {
		box += 0.10
		fmt.Println(box)
	}
}