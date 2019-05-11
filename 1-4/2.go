package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
)

func main() {
	var filename string
	const max_length int = 20
	fmt.Printf("Enter a file name: ")
	if _, err := fmt.Scan(&filename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	type Person struct {
		fname string
		lname string
	}

	var slice = make([]Person, 0)

	bytedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := string(bytedata)
	data = strings.TrimSpace(data)
	lines := strings.Split(data, "\n")
	for _, value := range lines {
		full_name := strings.Split(value, " ")
		fname := full_name[0]
		lname := full_name[1]
		if len(fname) > max_length {
			fmt.Println("String exceeds a limit")
			os.Exit(1)
		}
		if len(lname) > max_length {
			fmt.Println("String exceeds a limit")
			os.Exit(1)
		}
		person := Person{fname: fname, lname: lname}
		slice = append(slice, person)
	}

	for _, person := range slice {
		fmt.Println()
		fmt.Println("First name:", person.fname)
		fmt.Println("Last name:", person.lname)
	}
}
