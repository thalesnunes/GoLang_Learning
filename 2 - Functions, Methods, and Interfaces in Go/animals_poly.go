package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func makeAction(ani Animal, act string) {
	actions := map[string]func(){"eat": ani.Eat,
		"move":  ani.Move,
		"speak": ani.Speak}
	actions[act]()

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	animals := map[string]Animal{}
	var txt []string
	for {
		fmt.Println("If you want to exit, enter 'X'.")
		fmt.Printf("> ")
		scanner.Scan()
		if strings.ToLower(scanner.Text()) == "x" {
			fmt.Printf("Goodbye!")
			break
		}
		txt = strings.Fields(strings.ToLower(scanner.Text()))
		switch txt[0] {
		case "newanimal":
			switch txt[2] {
			case "cow":
				animals[txt[1]] = new(Cow)
			case "bird":
				animals[txt[1]] = new(Bird)
			case "snake":
				animals[txt[1]] = new(Snake)
			}
			fmt.Println("Created it!")
		case "query":
			switch txt[2] {
			case "eat":
				makeAction(animals[txt[1]], txt[2])
			case "move":
				makeAction(animals[txt[1]], txt[2])
			case "speak":
				makeAction(animals[txt[1]], txt[2])
			}
		default:
			fmt.Println("Enter a valid command.")
		}
	}
}
