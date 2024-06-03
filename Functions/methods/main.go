package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	status bool
}

func main() {
	fmt.Println("Methods in Golang")

	pranit := User{"Pranit", 22, "pranit@gmail.com", false}

	// fmt.Println("The user is", pranit)

	// fmt.Printf("The user is %+v\n", pranit)
	// fmt.Printf("The user name is %v and the email is %v\n", pranit.Name, pranit.Email)

	pranit.GetUserStatus()
	pranit.GetUserName()
	pranit.GetUserEmail()

}

func (u User) GetUserStatus() {
	fmt.Println("Is User active :", u.status)
}
func (u User) GetUserName() {
	fmt.Println("User Name :", u.Name)
}
func (u User) GetUserEmail() {
	fmt.Println("User Email :", u.Email)
}
