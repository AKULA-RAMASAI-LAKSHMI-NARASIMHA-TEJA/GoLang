package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	myName, _ := reader.ReadString('\n') // \n makes it reads until hitting new line
	fmt.Println(myName)
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')
	ageNum, _ := strconv.ParseFloat(strings.TrimSpace(age), 64)
	fmt.Println(ageNum)
}

/*
	// Scanln takes only first word as input
	var myStr string
	fmt.Scanln(&myStr)
	fmt.Println(myStr)

	var name string = "Teja"
	var a = fmt.Println(name) // a, _ to avoid error
	fmt.Println(a)
	// # command-line-arguments
	// .\readInput.go:9:10: multiple-value fmt.Println(name) (value of type (n int, err error)) in single-value context
*/
