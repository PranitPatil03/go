package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("User Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the you age:")

	age, _ := reader.ReadString('\n')
	fmt.Println("You age is - ", age)

}
