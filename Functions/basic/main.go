package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	greeter()

	result := adder(4, 3)
	fmt.Println("Sum of Two Numbers is", result)

	proResult := proAdder(4, 3, 3, 6, 2, 3, 5)
	fmt.Println("Sum of Two Numbers is", proResult)

	proResultWithMessage, myMessage := proAdderWithMessage(4, 3, 3, 6, 2, 3, 5)
	fmt.Println("Sum of Two Numbers is", proResultWithMessage)
	fmt.Println("Message is", myMessage)

}

func greeter() {
	fmt.Println("Hello, World!")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func proAdderWithMessage(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "`This is total sum`"
}
