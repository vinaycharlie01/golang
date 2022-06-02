package main

import "fmt"

type person struct {
	age int
}

func foo(a *person) func() *int {
	return func() *int {
		b := 10
		(*a).age = b
		return &a.age
	}

}

func main() {
	p1 := person{
		age: 21,
	}
	fmt.Println(p1)
	b := foo(&p1)
	c := &b
	fmt.Println("", *c)
}
