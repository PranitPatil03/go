package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["TS"] = "typescript"
	languages["PY"] = "python"

	fmt.Println("Map of languages is", languages)

	delete(languages, "PY")

	fmt.Println("Map of languages is", languages)

	for key, value := range languages {
		fmt.Printf("The Value of Key is %v, and the value is %v\n", key, value)
	}

}
