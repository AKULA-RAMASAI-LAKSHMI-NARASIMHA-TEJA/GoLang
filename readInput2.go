package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var userRating string

	// name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter name: ")
	name, _ := reader.ReadString('\n')

	// rating as input
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter user rating: ")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	fmt.Printf("Hello %v \nThanks for %v star rating. \n\nYour rating has recorded in our system at %v.", name, myNumRating, time.Now().Format(time.Stamp))

}
