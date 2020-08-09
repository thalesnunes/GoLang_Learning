package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Swap(sli []int, index int) { sli[index], sli[index+1] = sli[index+1], sli[index] }

func BubbleSort(sli []int) {

	swapped := true
	for swapped == true {
		swapped = false
		for now := 0; now < len(sli)-1; now++ {
			if sli[now] > sli[now+1] {
				Swap(sli, now)
				swapped = true
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("How many numbers do you want to sort?: ")
	scanner.Scan()
	numbInputs, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Printf("Type an integer.")
	}
	sli := make([]int, 0)
	var new int
	for i := 0; i < numbInputs; i++ {
		fmt.Printf("%dth number: ", i+1)
		scanner.Scan()
		new, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Type an integer.")
		}
		sli = append(sli, new)
	}
	fmt.Println("Slice before sorting:")
	fmt.Println(sli)
	BubbleSort(sli)
	fmt.Println("Slice after sorting:")
	fmt.Println(sli)
}
