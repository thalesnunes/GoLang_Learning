package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(accel, v0, s0 float64) func(float64) float64 {
	return func(time float64) float64 { return accel/2*math.Pow(time, 2) + v0*time + s0 }
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Acceleration value: ")
	scanner.Scan()
	accel, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Printf("Initial Velocity value: ")
	scanner.Scan()
	v0, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Printf("Initial Displacement value: ")
	scanner.Scan()
	s0, _ := strconv.ParseFloat(scanner.Text(), 64)
	fn := GenDisplaceFn(accel, v0, s0)
	var t float64
	for {
		fmt.Printf("Type the time 't' (if you want to exit, type X): ")
		scanner.Scan()
		if strings.ToLower(scanner.Text()) == "x" {
			break
		} else {
			t, _ = strconv.ParseFloat(scanner.Text(), 64)
			fmt.Printf("s(%.1f) = %.1f\n", t, fn(t))
		}
	}

}
