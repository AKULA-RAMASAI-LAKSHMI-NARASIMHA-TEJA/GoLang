package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["teja"] = 22
	ages["mahesh"] = 20
	ages["venkat"] = 21
	ages["abhi"] = 20
	ages["yash"] = 21

	fmt.Println(ages) // map[abhi:20 mahesh:20 teja:22 venkat:21 yash:21]

	delete(ages, "teja")
	fmt.Println(ages) // map[abhi:20 mahesh:20 venkat:21 yash:21]

}
