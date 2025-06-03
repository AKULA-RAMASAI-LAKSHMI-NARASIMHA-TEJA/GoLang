package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "working on files in go"
	file, err := os.Create("./myFile.txt")
	checkError(err)

	len, err := io.WriteString(file, content) // for writing io
	checkError(err)
	fmt.Println("File created successfully with Length:", len)
	readFile("./myFile.txt")

	defer file.Close() // for best practice

}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println(dataByte)         // [119 111 114 107 105 110 103 32 111 110 32 102 105 108 101 115 32 105 110 32 103 111]
	fmt.Println(string(dataByte)) // working on files in go
}

func checkError(err error) {
	if err != nil {
		panic(err) // panic will shut down the execution of the program and shows the error
	}
}
