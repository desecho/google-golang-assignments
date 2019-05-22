package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, i int) {
	currentElement := slice[i]
	nextElement := slice[i+1]
	slice[i+1] = currentElement
	slice[i] = nextElement
}

func BubbleSort(slice []int) {
	sliceLength := len(slice)
	for i, v := range slice {
		if i+1 == sliceLength {
			break
		}
		if v > slice[i+1] {
			Swap(slice, i)
			BubbleSort(slice)
		}
	}
}

func GetStringFromInput() string {
	var x string
	fmt.Printf("Enter ten integers separated by space: ")
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
	return slice
}

func main() {
	slice := GetSliceFromInput()
	BubbleSort(slice)
	fmt.Println("Sorted:", slice)
}
