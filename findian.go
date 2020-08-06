package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Type the string you want to parse: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.ToLower(scanner.Text())
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
