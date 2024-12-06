package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://facebook.com"

func main() {
	fmt.Println("web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// the type will return a pointer
	// this will be a reference of the original response not a copy of the response
	fmt.Printf("type of response %T\n", response)
	fmt.Printf("response value %v\n", response)

	// response body has to be closed
	// defer as this has to be last line to run
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
