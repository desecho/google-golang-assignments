// Program with a race condition. It can print either 0 or 1.

package main

import (
	"fmt"
)

func AddOneToX(x *int) {
	*x += 1
}

func PrintX(x *int) {
	fmt.Println(*x)
}

func main() {
	var x int
	fmt.Println("Press enter to quit")
	for i := 0; i < 20; i++ {
		go AddOneToX(&x)
		go PrintX(&x)
	}
	fmt.Scanln()
}

/* Race condition occurs when two or more goroutines are executed at the same time.
   It happens because of the interleavings which means that instructions do not execute
   in a deterministic order, sometimes instruction 1 of one goroutine can be executed
   before instruction 1 of another goroutine, etc. */
