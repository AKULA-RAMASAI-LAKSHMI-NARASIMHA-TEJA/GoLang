package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://example.com"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", response) // *http.Response

	data, err := ioutil.ReadAll(response.Body) // returns data in bytes
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) // html

	defer response.Body.Close() // best practice
}
