package main

import (
	"fmt"
	"math"
	"os"
)

func GetInputs() (float64, float64, float64, float64) {
	var a, v0, s0, t float64
	fmt.Printf("Enter acceleration: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Enter initial velocity: ")
	if _, err := fmt.Scan(&v0); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Enter initial displacement: ")
	if _, err := fmt.Scan(&s0); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Enter time: ")
	if _, err := fmt.Scan(&t); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return a, v0, s0, t
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 { return 0.5*a*math.Pow(t, 2) + v0*t + s0 }
}

func main() {
	a, v0, s0, t := GetInputs()
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement: %f\n", fn(t))
}
