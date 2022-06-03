package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

type giomatric interface {
	shape
	object
	getcolor() string
}

type cune struct {
	edge  float64
	color string
}

func (c cune) area() float64 {
	return 6 * (c.edge * c.edge)

}
func (c cune) volume() float64 {
	return math.Pow(c.edge, 3)

}

func (c cune) getcolor() string {
	return c.color

}
func measure(g giomatric) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v

}

func main() {
	c := cune{edge: 2}
	a, v := measure(c)
	fmt.Println("", a, v)

}
