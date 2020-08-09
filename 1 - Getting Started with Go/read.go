package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	type Person struct {
		fname string
		lname string
	}

	fmt.Printf("Name of the text file: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()
	arq, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading the file")
		defer arq.Close()
	}
	scanner2 := bufio.NewScanner(arq)
	sli := make([]Person, 0)
	words := make([]string, 2)
	var newPer Person
	for scanner2.Scan() {
		words = strings.Fields(scanner2.Text())
		newPer = Person{fname: words[0], lname: words[1]}
		sli = append(sli, newPer)
	}
	for _, val := range sli {
		fmt.Println(val.fname, val.lname)
	}

}
