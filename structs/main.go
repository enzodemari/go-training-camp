package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	contactInfo
}

func main() {
	//alex := Person{firstName: "Alex", lastName: "Anderson"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@mail.com",
			zipCode: 30987,
		},
		contactInfo: contactInfo{
			email:   "jim2.mail.com",
			zipCode: 45788,
		},
	}

	fmt.Printf("\n%+v", jim)
}
