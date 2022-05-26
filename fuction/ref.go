package main

import "fmt"

func heloo(a []string) {
	a[0] = "vinay"

}
func main() {
	a := make([]string, 1, 25)
	fmt.Println(a)
	heloo(a)
	fmt.Println(a)

}
