package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Switch in Golang")

	fmt.Println("Enter a number from 1 to 6")

	number := bufio.NewReader(os.Stdin)

	userNum, _ := number.ReadString('\n')

	fmt.Printf("%T\n", userNum)

	switch userNum {
	case "1":
		fmt.Println("You entered 1")
	case "2":
		fmt.Println("You entered 2")
	case "3":
		fmt.Println("You entered 3")
	case "4":
		fmt.Println("You entered 4")
	case "5":
		fmt.Println("You entered 5")
	case "6":
		fmt.Println("You entered 6")
	default:
		fmt.Println("You entered invalid number")
	}

}
