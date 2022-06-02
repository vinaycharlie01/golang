package main

import "fmt"

type car struct {
	brand  string
	prince int
}

func chamgecar(c car, newbrand string, newprice int) {
	c.brand = newbrand
	c.prince = newprice

}

func (c car) chamgecar1(newbrand string, newprice int) {
	c.brand = newbrand
	c.prince = newprice

}
func (c *car) chamgecar2(newbrand string, newprice int) {
	(*c).brand = newbrand
	(*c).prince = newprice

}

func main() {
	cars := car{
		brand:  "audi",
		prince: 2000,
	}
	chamgecar(cars, "aude", 2000)
	fmt.Println(cars)
	cars.chamgecar1("mpower", 3000)
	fmt.Println(cars)
	(&cars).chamgecar2("cinay", 209993)
	fmt.Println(cars)
	var a *car
	a = &cars
	fmt.Println(*a)

}
