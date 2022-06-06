package main

import "fmt"

func factorial(a int, c chan int) {
	f := 1
	for i := 1; i <= a; i++ {
		f *= i
	}
	c <- f

}

func main() {
	ch := make(chan int)
	go factorial(5, ch)
	f := <-ch
	fmt.Println("", f)

}

