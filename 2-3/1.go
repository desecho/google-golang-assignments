package main

import (
	"fmt"
	"os"
)

type Animal struct {
	food     string
	movement string
	sound    string
}

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Move() string {
	return animal.movement
}

func (animal Animal) Speak() string {
	return animal.sound
}

func GetDataFromInput() (string, string) {
	fmt.Printf("> ")
	var animal, action string
	if _, err := fmt.Scan(&animal, &action); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return animal, action
}

func main() {
	cow := Animal{food: "grass", movement: "walk", sound: "moo"}
	bird := Animal{food: "worms", movement: "fly", sound: "peep"}
	snake := Animal{food: "mice", movement: "slither", sound: "hsss"}
	animals := map[string]Animal{"cow": cow, "bird": bird, "snake": snake}
	for {
		animalChosen, action := GetDataFromInput()
		animal := animals[animalChosen]
		switch action {
		case "eat":
			fmt.Println(animal.Eat())
		case "move":
			fmt.Println(animal.Move())
		case "speak":
			fmt.Println(animal.Speak())
		}
	}
}
