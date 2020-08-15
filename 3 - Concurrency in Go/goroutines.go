package main

import (
	"fmt"
	"sync"
)

func SumOne(x *int, wg *sync.WaitGroup) {
	*x = *x + 1
	defer wg.Done()
}

func MultTwo(x *int, wg *sync.WaitGroup) {
	*x = *x * 2
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	x := 3
	wg.Add(1)
	go SumOne(&x, &wg)
	wg.Wait()
	wg.Add(1)
	go MultTwo(&x, &wg)
	wg.Wait()
	fmt.Println(x)
}
