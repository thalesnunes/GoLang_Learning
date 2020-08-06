package main

import "fmt"

func main() {

	fmt.Printf("What number do you want to truncate?: ")
	var floa float32
	fmt.Scan(&floa)
	fmt.Printf("%d", int(floa))
}
