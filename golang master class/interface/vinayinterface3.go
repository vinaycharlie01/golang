package main

import "fmt"

type person struct {
	name string
	age  int
}
type alldel interface {
	hello() int
	hello1() string
}

func (p person) hello() int {
	return p.age + 10

}
func (p person) hello1() string {
	return p.name + "hello"

}

func hello(a alldel) (int, string) {
	b := a.hello()
	c := a.hello1()
	return b, c

}

func main() {
	p1 := person{
		name: "Vinay",
		age:  10,
	}
	a, b := hello(p1)
	fmt.Println("", a, b)
}
