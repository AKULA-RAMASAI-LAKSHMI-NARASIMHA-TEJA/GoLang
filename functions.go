package main

import "fmt"

func main() {
	greet()
	fmt.Println(mul(10, 10))
	fmt.Println(add(10, 10))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(return2val())
}

func greet() {
	fmt.Println("Hiee, How r u ?")
}

func mul(a int, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func sum(values ...int) int {
	ans := 0
	for i := range values {
		ans += values[i]
	}
	return ans
}

func return2val() (int, string) {
	return 39, "Teja"
}

/*
Hiee, How r u ?
100
20
15
39 Teja
*/
