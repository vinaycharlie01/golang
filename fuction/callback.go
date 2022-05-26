package main

import "fmt"

func hello(a []int, b func(int) bool) []int {
	xs := []int{}
	for _, v := range a {
		if b(v) {
			xs = append(xs, v)
		}
	}
	return xs

}
func main() {
	xs := hello([]int{10, 20, 30}, func(i int) bool {
		return i > 1
	})
	fmt.Println((xs))

}
