package main

import "fmt"

func main() {
	fmt.Println("Defer in golang")
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("welcome")
	Defer()
}

func Defer() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
