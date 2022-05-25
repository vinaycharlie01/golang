package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)

}

func main() {

	// alex := person{"Alex", "Anderson"} 1
	// alex := person{firstName: "Alex", lastName: "Anderson"} 2
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "James" 3
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			pincode: 563101,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("KALYAN")
	jim.print()

}
