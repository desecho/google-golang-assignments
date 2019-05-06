package main

import "fmt"
import "strconv"
import "sort"
func main() {
	var input string
	lst := make([]int, 0, 3)
	for {
		fmt.Println("please type a integer to contine or X to exit")
		fmt.Scan(&input)
		if input == "X"{
			break
		}
		i, err := strconv.Atoi(input)
	    if err == nil {
					lst = append(lst, i)
					sort.Ints(lst)
					fmt.Println("sorted current array: ", lst)

	    }

	}


}
