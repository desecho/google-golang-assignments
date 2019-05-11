package main

import "fmt"
import "os"

func main() {
    var x float32
    fmt.Printf("Enter a floating point number: ")
    if _, err := fmt.Scan(&x); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    y := int(x)
    fmt.Printf("%d\n", y)
}

