package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var x string
	var slice = make([]int, 0, 3)

	for {
		fmt.Printf("Enter an integer: ")
		if _, err := fmt.Scan(&x); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if x == "X" {
			break
		}

		y, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		slice = append(slice, y)
		sort.Ints(slice)
		fmt.Println("Sorted: ", slice)
	}
}
