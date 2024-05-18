package main

import "fmt"

func main() {

	fmt.Println("The Pointers in Golang")

	var ptr *int

	fmt.Println("The Value of ptr is", ptr)

	myAge := 22

	var newPtr = &myAge

	fmt.Println("The Value of ptr is", newPtr)
	fmt.Println("The Value of ptr is", *newPtr)

	*newPtr = *newPtr + 2

	fmt.Println("The New Value of ptr is", myAge)

}
