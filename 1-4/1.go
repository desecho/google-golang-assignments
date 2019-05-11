package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"encoding/json"
)

func main() {
	var name, address string
	var err error

	fmt.Printf("Enter a name: ")
	in := bufio.NewReader(os.Stdin)
	name, err = in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	name = strings.TrimSpace(name)
	fmt.Printf("Enter an address: ")

	address, err = in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	address = strings.TrimSpace(address)

	person := map[string]string {"name": name, "address": address}
	json, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("JSON:", string(json))
}
