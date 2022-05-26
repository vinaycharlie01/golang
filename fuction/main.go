package main

import "fmt"

func foo(a ...int) int {
	val := 0
	for _, v := range a {
		val += v

	}
	return val
}
func foo1(b []int) int {
	val := 0
	for _, v := range b {
		val += v

	}
	return val
}

func main() {
	a := []int{10, 20, 30, 30, 50, 60}
	b := foo(a...)
	d := foo1(a)
	fmt.Println(b, d)
}

