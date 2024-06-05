package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://x.com/home"

func main() {
	fmt.Println("Web Request handing in Golang")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	dataBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println("Url Content is - ", content)

}
