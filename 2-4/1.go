package main

import (
	"fmt"
	"os"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (_ *Cow) Eat() {
	fmt.Println("grass")
}

func (_ *Cow) Move() {
	fmt.Println("walk")
}

func (_ *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (_ Bird) Eat() {
	fmt.Println("worms")
}

func (_ Bird) Move() {
	fmt.Println("fly")
}

func (_ Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (_ Snake) Eat() {
	fmt.Println("mice")
}

func (_ Snake) Move() {
	fmt.Println("slither")
}

func (_ Snake) Speak() {
	fmt.Println("hsss")
}

func GetDataFromInput() (string, string, string) {
	fmt.Printf("> ")
	var command, name, thirdValue string
	if _, err := fmt.Scan(&command, &name, &thirdValue); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return command, name, thirdValue
}

func main() {
	animals := map[string]Animal{}
	for {
		command, name, thirdValue := GetDataFromInput()
		switch command {
		case "newanimal":
			animalType := thirdValue
			switch animalType {
			case "cow":
				var cow *Cow
				animals[name] = cow
				fmt.Println("Created it!")
			case "bird":
				var bird *Bird
				animals[name] = bird
				fmt.Println("Created it!")
			case "snake":
				var snake *Snake
				animals[name] = snake
				fmt.Println("Created it!")
			}
		case "query":
			action := thirdValue
			animal := animals[name]
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		}
	}
}
