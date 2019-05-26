package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
Let us assume the following formula for displacement s as a function of time t, acceleration a,
initial velocity vo, and initial displacement so.

s =½at^2 + vt + s

Write a program which first prompts the user to enter values for acceleration, initial velocity,
and initial displacement. Then the program should prompt the user to enter a value for time and
the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time, assuming the given
values acceleration, initial velocity, and initial displacement. The function returned by
GenDisplaceFn() should take one float64 argument t, representing time, and return one float64
argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial
velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement
to call GenDisplaceFn() to generate a function fn which will compute displacement as a function
of time.

fn := GenDisplaceFn(10, 2, 1)
*/

func main() {
	var vStr string
	var aStr string
	var sStr string

	fmt.Print("Input initial acceleration: ")
	fmt.Scan(&aStr)
	fmt.Print("Input initial velocity: ")
	fmt.Scan(&vStr)
	fmt.Print("Input initial displacement: ")
	fmt.Scan(&sStr)

	a, _ := strconv.ParseFloat(aStr, 64)
	v, _ := strconv.ParseFloat(vStr, 64)
	s, _ := strconv.ParseFloat(sStr, 64)

	displacement := GenDisplaceFn(a, v, s)

	var tStr string
	for {
		fmt.Print("What time (X/x to exit)? ")
		fmt.Scan(&tStr)
		if strings.ToLower(tStr) == "x" {
			break
		}
		t, _ := strconv.ParseFloat(tStr, 64)
		fmt.Println(displacement(t))
	}
}

func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return a*math.Pow(t, 2) + v*t + s
	}
}
