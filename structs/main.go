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
}

func main() {
	// jp := person{firstName: "JyothiPrasad"}
	var jp person
	jp.firstName = "JyothiPrasad"
	jp.contact.email = "jyothiprasadu@gmail.com"
	fmt.Println(jp)
	fmt.Printf("%+v", jp)
}
