package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	status bool
}

func main() {
	fmt.Println("Structs in Golang")

	pranit := User{"Pranit", 22, "pranit@gmail.com", true}

	fmt.Println("The user is", pranit)

	fmt.Printf("The user is %+v\n", pranit)
	fmt.Printf("The user name is %v and the email is %v\n", pranit.Name, pranit.Email)

}
