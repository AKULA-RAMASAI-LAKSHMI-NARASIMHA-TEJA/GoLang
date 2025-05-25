package main

import "fmt"

func main() {
	var batman string = "I am a batman"
	fmt.Println(batman)

	var superman = "I am a superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	thorRating := 3.0
	// v is for value and T is for type
	fmt.Printf("%v %T \n", thorRating, thorRating)

	var a, b string = "I am A", "I am B"
	fmt.Println(a, b)

	//This type of declaration is not objects it is known as Standalone variables in Go
	var (
		name = "Teja"
		age  = 22
		role = "SDE"
	)
	fmt.Println(name, age, role)
}
