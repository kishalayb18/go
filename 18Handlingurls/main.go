package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=go&class=hy&name=kishalay"

func main() {
	fmt.Println("go handling urls")
	fmt.Println(myurl)

	// parsing
	fmt.Println("\nparsing url")
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println("\n")

	// this will be key value pairs
	// both key value will be type string
	fmt.Println("query parameter")
	queryParams := result.Query()
	fmt.Printf("the type of query parameter %T \n", queryParams)

	// passing the key of the url query to get the value
	fmt.Println(queryParams["coursename"])
	fmt.Println(queryParams["name"])

	// for loop to get values from queryParams
	for _, val := range queryParams {
		fmt.Println("param value ", val)
	}

	// build a new url from the datas
	urlDetails := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=kish",
	}

	newurl := urlDetails.String()
	fmt.Println("new url ", newurl)

}
