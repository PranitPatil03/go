package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("User Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the you age:")

	age, _ := reader.ReadString('\n')
	fmt.Println("You age is - ", age)

	ageplus1, err := strconv.ParseFloat(strings.TrimSpace(age), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your age - ", ageplus1+1)
	}

}
