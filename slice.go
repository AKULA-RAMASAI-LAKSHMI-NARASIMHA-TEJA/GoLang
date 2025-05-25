package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice) // [1 2 3 4 5]

	slice = append(slice, 6)
	fmt.Println(slice) // [1 2 3 4 5 6]

	slice = append(slice[1 : len(slice)-1])
	fmt.Println(slice)

	colors := make([]string, 3, 3)
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"
	fmt.Println(colors) // [red blue green]
	colors = append(colors, "yellow", "orange")
	fmt.Println(colors) // [red blue green yellow orange]

	slice = append(slice, 1, 2, 2, 3)
	isSorted := sort.IntsAreSorted(slice)
	fmt.Println(isSorted) // false
	sort.Ints(slice)
	fmt.Println(slice) // [1 2 2 2 3 3 4 5]
}
