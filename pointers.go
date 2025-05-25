package main

import "fmt"

func main() {
	var num int = 10
	numRef := &num // numRef stores the address of num variable

	fmt.Println(num, numRef) // 10 0xc00000a0f8

	num *= 2

	// '*' is used to access the value stored in that address
	fmt.Println(num, *numRef) // 20 20
}
