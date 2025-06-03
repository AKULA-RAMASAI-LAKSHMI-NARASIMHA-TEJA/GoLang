package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=golang&paymentid=gvujhi9"

func main() {
	fmt.Println(myUrl) // https://lco.dev:3000/learn?coursename=golang&paymentid=gvujhi9

	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)          // https://lco.dev:3000/learn?coursename=golang&paymentid=gvujhi9
	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // lco.deev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.Port())   // 3000
	fmt.Println(result.RawQuery) // coursename=golang&paymentid=gvujhi9

	qparams := result.Query()
	fmt.Println(qparams) // map[coursename:[golang] paymentid:[gvujhi9]]

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=teja",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl) // https://lco.dev/tutcss
}
