package main

import (
	"fmt"
)

//variable
var a int
var b string = "Vinau"
var c = true

var e, f, g string
var a1, b1, c1 int = 10, 20, 30
var (
	a2     = 30
	b2 int = 30
)

func main() {
	e, f, g = "vinay", "kumar", "va"
	a = 230
	fmt.Println("this is your value: ", a)
	fmt.Printf("%v", b)
	fmt.Println(c)
	d := 20
	fmt.Println("t", d)
	fmt.Println("your name :", e, f, g)
	e1, g1 := 10, 20
	fmt.Println(e1, g1)
	fmt.Println(a1, b1, c1)
	fmt.Println("", a2, b2)
}
