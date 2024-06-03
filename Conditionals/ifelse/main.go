package main

import "fmt"

func main() {
	fmt.Println("If-Else in Golang")

	loginCount := 30

	var result string

	if loginCount < 10 {
		result = "Low"
	} else if loginCount > 10 {
		result = "High"
	} else {
		result = "Medium"
	}

	fmt.Println("The Result is", result)

	if 5%2 == 0 {
		fmt.Println("5 is even")
	} else {
		fmt.Println("5 is odd")
	}

	if num := 3; num > 10 {
		fmt.Println("num is greater than 10")
	} else {
		fmt.Println("num is less than 10")
	}

}
