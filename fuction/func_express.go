package main

import "fmt"

func main() {
	a := func(a []int) int {
		val := 0
		for _, v := range a {
			val += v

		}
		return val
	}
	b := a([]int{10, 20, 30, 30})
	fmt.Println(b)

}
