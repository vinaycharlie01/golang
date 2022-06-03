package main

import "fmt"

type a1 interface {
	hello() string
}
type a2 interface {
	what() int
}
type all interface {
	a1
	a2
	how() float64
}

type person struct {
	name  string
	age   int
	hight float64
}

func (a person) hello() string {
	return a.name

}
func (b person) what() int {
	return b.age

}
func (c person) how() float64 {
	return c.hight

}
func hello(c person) (string, int, float64) {
	w := c.hello()
	x := c.what()
	y := c.how()
	return w, x, y

}

func main() {
	q := person{
		name:  "Vinay",
		age:   21,
		hight: 5.59,
	}
	j, k, l := hello(q)
	fmt.Println("", j, k, l)

}
