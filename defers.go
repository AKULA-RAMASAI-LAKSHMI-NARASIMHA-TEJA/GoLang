package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")

	/* Defer key word waits the execution until the nearest function returns and it follows LIFO order
	hello
	two
	one
	world
	*/

	defer myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // prints 0 to 4 in reverse order
	}
}
