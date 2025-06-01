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

	// just checking random stuff
	scores := make([]int, 4)
	scores[0] = 7
	scores[1] = 4
	scores[2] = 8
	scores[3] = 6
	fmt.Println("scores address: ", &scores[0])
	ptr := &scores[0]

	scores = append(scores, 5, 1, 2)
	fmt.Println("scores address after appendinng elements: ", &scores[0])
	*ptr *= 2
	fmt.Println(*ptr, scores[0])
}
