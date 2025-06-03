package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	Teja := User{"Teja", "teja@gmail.com", true, 22}
	fmt.Println(Teja) // {Teja teja@gmail.com true 22}
	Teja.getStatus()
	Teja.newEmail()
	fmt.Println(Teja) // {Teja teja@gmail.com true 22}

}

func (u User) getStatus() {
	fmt.Println(u.Status) //true
}

func (u User) newEmail() {
	u.Email = "tej@gmail.com"
	fmt.Println(u) // {Teja tej@gmail.com true 22}
	// here the mail changes but it doesn't change in original object it only passes the copy of original object
}
