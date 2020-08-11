package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (ani Animal) Eat() {
	fmt.Println(ani.food)
}

func (ani Animal) Move() {
	fmt.Println(ani.locomotion)
}

func (ani Animal) Speak() {
	fmt.Println(ani.noise)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
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
		case "cow":
			switch txt[1] {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}
		case "bird":
			switch txt[1] {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}
		case "snake":
			switch txt[1] {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}
		default:
			fmt.Println("Enter a valid animal.")
		}
	}
}
