package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")

	var marks [5]int

	marks[0] = 54
	marks[1] = 65
	marks[2] = 65
	marks[3] = 34
	marks[4] = 67

	fmt.Println("The Marks Array is", marks)
	fmt.Println("The Marks Array is", len(marks))

	var todos = [3]string{"Project", "GYM", "Work"}
	fmt.Println("The Marks Array is", todos)
	fmt.Println("The Marks Array is", len(todos))

}
