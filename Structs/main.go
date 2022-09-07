package main

import (
	"fmt"
)

type contactInfo struct{
	email string
	zipcode int
}

type Person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	user1 := Person{firstName: "shashank", lastName: "thakur"}
	fmt.Println((user1))

	var user2 Person
	user2.firstName = "Veer"
	user2.lastName = "Thakur"
	// fmt.Printf("%+v", user2)

	user3 := Person{
		firstName: "John",
		lastName: "Doe",
		contactInfo: contactInfo{
			email: "abc@xyz.com",
			zipcode: 123456,
		},
	}
	// fmt.Printf("%+v", user3)

	userPointer := &user3
	userPointer.updateName("Max")
	fmt.Printf("%+v", user3)
}

func (p *Person) updateName(newFirstName string){
	(*p).firstName = newFirstName
}