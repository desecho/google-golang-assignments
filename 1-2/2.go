package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var x string
	fmt.Printf("Enter a string: ")
	in := bufio.NewReader(os.Stdin)

	x, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	x = strings.ToLower(x)
	x = strings.TrimSpace(x)
	first_char := string(x[0])
	last_char := string(x[len(x)-1])
	if first_char == "i" && last_char == "n" && strings.Contains(x, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
