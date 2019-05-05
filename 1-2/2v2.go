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
	if strings.HasPrefix(x, "i") && strings.Contains(x, "a") && strings.HasSuffix(x, "n") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
