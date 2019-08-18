package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for true {
		fmt.Println("Enter animal name space seperated by action")
		fmt.Println(">")

		var inStr string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			inStr = scanner.Text()
		}
		input := strings.Split(inStr, " ")

		animal := strings.TrimSpace(input[0])
		action := strings.TrimSpace(input[1])

		cow := Animal{"grass", "walk", "moo"}
		bird := Animal{"worms", "fly", "peep"}
		snake := Animal{"mice", "slither", " hsss"}

		switch animal {
		case "cow":
			switch action {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			default:
				fmt.Println("Incorrect action provided. Please provide eat, move or speak")

			}
		case "bird":
			switch action {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			default:
				fmt.Println("Incorrect action provided. Please provide eat, move or speak")
			case "speak":
				bird.Speak()
			}
		case "snake":
			switch action {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			default:
				fmt.Println("Incorrect action provided. Please provide eat, move or speak")
			}
		default:
			fmt.Println("Incorrect Animal provided. Please provide cow, bird or snake")
		}

	}
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}
