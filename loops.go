package main

import "fmt"

func main() {
	start := 1
	things := []string{"bat", "ball", "book", "pen", "laptop"}

	fmt.Println(things)

	// printing 1 to 10 using for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i+start, " ")
	}
	fmt.Println()

	for i := 0; i < len(things); i++ {
		fmt.Print(things[i], " ")
	}
	fmt.Println()

	for i := range things {
		fmt.Print(things[i], " ")
	}
	fmt.Println()

	for start < 100 {
		start += start
		if start == 32 {
			goto label
		}
		fmt.Println("start is now at:", start)
	}

label:
	fmt.Println("goto statement")
	fmt.Println("2nd line in label")
}

/*
[bat ball book pen laptop]
1 2 3 4 5 6 7 8 9 10
bat ball book pen laptop
bat ball book pen laptop
start is now at: 2
start is now at: 4
start is now at: 8
start is now at: 16
goto statement
2nd line in label
*/
