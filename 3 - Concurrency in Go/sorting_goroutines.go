package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

func PartitionArrayFour(sli []int) ([]int, []int, []int, []int) {
	sizeFourth := int(len(sli) / 4)
	return sli[:sizeFourth], sli[sizeFourth : sizeFourth*2], sli[sizeFourth*2 : sizeFourth*3], sli[sizeFourth*3:]
}

func PrintSort(sli []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(sli)
	sort.Ints(sli)
}

func main() {

	var wg sync.WaitGroup
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
	sli1, sli2, sli3, sli4 := PartitionArrayFour(sli)
	wg.Add(4)
	go PrintSort(sli1, &wg)
	go PrintSort(sli2, &wg)
	go PrintSort(sli3, &wg)
	go PrintSort(sli4, &wg)
	wg.Wait()
	sliceFinal := sli1[:]
	sliceFinal = append(sliceFinal, sli2...)
	sliceFinal = append(sliceFinal, sli3...)
	sliceFinal = append(sliceFinal, sli4...)
	sort.Ints(sliceFinal)
	fmt.Println(sliceFinal)
}
