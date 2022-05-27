package main

import "fmt"

type person struct {
	name string
}

type viii struct {
	last string
}

type vinay interface {
	hello() string
}

func (a person) hello() string {
	return a.name
}

func (b viii) hello() string {
	return b.last

}

func info(a vinay) {
	fmt.Println(a.hello())

}

func main() {
	p1 := person{
		name: "Vinay",
	}
	info(p1)
}
