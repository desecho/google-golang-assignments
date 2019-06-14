package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetStringFromInput() string {
	var x string
	fmt.Printf("Enter integers separated by space: ")
	in := bufio.NewReader(os.Stdin)

	x, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	x = strings.TrimSpace(x)
	return x
}

func GetSliceFromInput() []int {
	x := GetStringFromInput()
	s := strings.Split(x, " ")
	slice := make([]int, 0)
	for _, element := range s {
		y, err := strconv.Atoi(element)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		slice = append(slice, y)
	}
	if len(slice) < 4 {
		fmt.Println("You need to enter at least 4 integers")
		os.Exit(1)
	}
	return slice
}

func Sort(slice []int, c chan []int) {
	fmt.Println(slice)
	sort.Ints(slice)
	c <- slice
}

func main() {
	slice := GetSliceFromInput()
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	s3 := make([]int, 0)
	s4 := make([]int, 0)
	c := make(chan []int)

	for {
		if len(slice) < 1 {
			break
		}
		s1 = append(s1, slice[0])
		slice = slice[1:len(slice)]

		if len(slice) < 1 {
			break
		}
		s2 = append(s2, slice[0])
		slice = slice[1:len(slice)]

		if len(slice) < 1 {
			break
		}
		s3 = append(s3, slice[0])
		slice = slice[1:len(slice)]

		if len(slice) < 1 {
			break
		}
		s4 = append(s4, slice[0])
		slice = slice[1:len(slice)]
	}

	go Sort(s1, c)
	go Sort(s2, c)
	go Sort(s3, c)
	go Sort(s4, c)
	s1 = <-c
	s2 = <-c
	s3 = <-c
	s4 = <-c
	slice = append(s1, s2...)
	slice = append(slice, s3...)
	slice = append(slice, s4...)
	sort.Ints(slice)
	fmt.Println()
	fmt.Println(slice)
}
