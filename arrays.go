package main

import "fmt"

func main() {

	// syntax: var arr_name [size]type
	var nums [3]string
	nums[0] = "one"
	nums[1] = "two"
	nums[2] = "three"
	fmt.Println(nums) // [one two three]

	var colors = [4]string{"red", "blue", "green", "orange"}
	fmt.Println(colors)      // [red blue green orange]
	fmt.Println(len(colors)) // 4
}
