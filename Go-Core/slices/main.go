package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices in Golang")

	var myMarks = []int{34, 65, 43, 453, 323}

	fmt.Println("My marks are -", myMarks)
	fmt.Printf("The Type of Marks is %T\n", myMarks)

	myMarks = append(myMarks, 12, 54)
	fmt.Println("My marks are -", myMarks)

	myMarks = append(myMarks[1:4])
	fmt.Println("My marks are -", myMarks)

	myMarks = append(myMarks[:4])
	fmt.Println("My marks are -", myMarks)

	highMarks := make([]int, 5)

	highMarks[0] = 343
	highMarks[1] = 454
	highMarks[2] = 6732
	highMarks[3] = 3343
	highMarks[4] = 345435

	// fmt.Println("High Marks are ", highMarks)

	// highMarks = append(highMarks, 545, 45, 43, 23)

	// fmt.Println("High Marks are ", highMarks)

	// sort.Ints(highMarks)
	// fmt.Println("High Marks are ", highMarks)

	// fmt.Println(sort.IntsAreSorted(highMarks))

	var courses = []string{"reactjs", "nodejs", "next", "javascript"}

	fmt.Println(courses)

	index := 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
