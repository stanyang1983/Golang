package main

import "fmt"

func main() {
	var x int = 0
	for x < 3 {
		fmt.Println(x)
		x++
	}

	var y int
	for y = 0; y < 3; y++ {
		fmt.Println(y)
	}

}
