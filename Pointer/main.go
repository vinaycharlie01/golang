package main

import "fmt"

func foo(a *int) *int {
	*a = 39
	return a

}

func main() {
	b := 10
	d := foo(&b)
	*d = 30
	fmt.Println(*d)

	// a := 10
	// b := &a
	// *b = 20

}
