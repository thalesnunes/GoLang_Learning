package main

import "fmt"

func main() {
	var z int = 10
	var y *int
	y = &z
	fmt.Print(*y)
}
