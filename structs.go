package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	teja := User{"Teja", "teja@gmail.com", 22}
	fmt.Println(teja)                            // {Teja teja@gmail.com 22}
	fmt.Printf("%v\n", teja)                     // {Teja teja@gmail.com 22}
	fmt.Printf("%+v\n", teja)                    // {Name:Teja Email:teja@gmail.com Age:22}
	fmt.Println(teja.Name, teja.Age, teja.Email) // Teja 22 teja@gmail.com

	sai := new(User)
	sai.Name = "Sai"
	sai.Age = 21
	fmt.Println(sai) // &{Sai  21}
	sai.Email = "sai@gmail.com"
	fmt.Println(sai)        // &{Sai sai@gmail.com 21}
	fmt.Printf("%v\n", sai) // &{Sai sai@gmail.com 21}

	sam := &User{"Sam", "sam@gmail.com", 30}
	fmt.Println(sam) // &{Sam sam@gmail.com 30}
}
