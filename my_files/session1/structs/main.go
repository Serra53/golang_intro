package main

import "fmt"

//structs are like dictionaries in Python or JSON in JavaScript

type person struct {
	firstName string
	lastName  string
	age       int64
	isSmoker  bool
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	person1 := person{
		firstName: "Jim",
		lastName:  "Carrey",
		age:       36,
		isSmoker:  true,
		contact: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}

	person1.print()
	//give me the memory adress of the value this variable is pointing at
	person1Pointer := &person1
	fmt.Println(person1Pointer)
	person1Pointer.updateFirstName("Michael")
	person1.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

//In the code below
//*person is a type of pointer to person
//*pointerToPerson is the actual memory value
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
