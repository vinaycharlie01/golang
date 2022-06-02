package main

import "fmt"

func foo(a *int) *int {
	*a = 20 * 28
	return a

}

func main() {

	a := 10
	b := &a
	d := foo(b)
	fmt.Println("", d)

	val := 20
	val1 := &val
	val2 := &val1
	fmt.Println("", *val1 == **val2)
	var p2 *int
	a = 20
	p2 = &a
	fmt.Println("", *p2)

}
