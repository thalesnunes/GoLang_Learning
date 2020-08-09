package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	fmt.Printf("What is your name? ")
	var name string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	fmt.Printf("What is your address? ")
	var address string
	scanner.Scan()
	address = scanner.Text()
	jso := map[string]string{"name": name, "address": address}
	prin, err := json.Marshal(jso)
	if err != nil {
		fmt.Println("Encoding Failed!")
	} else {
		fmt.Println(string(prin))
	}

}
