package main

import "fmt"

func heloo() {
	fmt.Println("Hello vinay")

}

func girl() {
	fmt.Println("i love you")

}

func main() {
	defer heloo()
	girl()

}
