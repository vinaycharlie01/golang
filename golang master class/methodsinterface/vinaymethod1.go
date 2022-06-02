package main

import "fmt"

type names []string

func (n names) print() {
	for _, v := range n {
		fmt.Println(v)
	}

}

func main() {
	p1 := names{"vinay", "kumar", "kalyan"}
	p1.print()
	names.print(p1)

}
