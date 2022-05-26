package main

import "fmt"

func fooo(x int) int {
	if x == 0 {
		return 1
	}
	return x * fooo(x-1)

}

func main() {
	fmt.Println(fooo(4))

}
