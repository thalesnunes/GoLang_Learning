package main

import "fmt"

func main() {

	var x = "h"
	var br = false

	for br == false {
		switch x {
		case "x":
			fmt.Printf("right1\n")
			x = "c"
		case "h":
			fmt.Printf("right2\n")
			x = "x"
		default:
			fmt.Printf("done")
			br = true
		}
	}
}
