package main

import "fmt"

func main() {
	fmt.Println("Loops in Golong")

	marks := []int{43, 54, 54, 65, 43, 54}

	fmt.Println("Marks", marks)

	// for i := 0; i < len(marks); i++ {
	// 	fmt.Print(marks[i], "  \n")
	// }

	// for i := range marks {
	// 	fmt.Print(marks[i], " ")
	// }
		
	for i, marks := range marks {
		fmt.Printf("Marks[%d] = %d\n", i, marks)
	}

}
