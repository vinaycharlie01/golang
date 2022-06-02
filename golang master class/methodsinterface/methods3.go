package main

import "fmt"

type person struct {
	names string
	age   int
}

func carchane(a person, namesp1 string, agep1 int) {
	a.names = namesp1
	a.age = agep1

}

func (c person) carchane1(namesp1 string, agep1 int) {
	c.names = namesp1
	c.age = agep1

}
func (c *person) carchane2(namesp1 string, agep1 int) {
	(*c).names = namesp1
	(*c).age = agep1

}

func main() {
	p1 := person{
		names: "vinay",
		age:   20,
	}
	carchane(p1, "adarsh", 19)
	fmt.Println(p1)
	p1.carchane1("vinay", 29)
	fmt.Println(p1)
	(&p1).carchane2("vinay", 29)
	fmt.Println(p1)
	var a *person
	a = &p1
	fmt.Println("", a)
	a.carchane2("vinau", 123)
	fmt.Println(*a)

}
